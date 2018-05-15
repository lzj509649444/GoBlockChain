# 创建coinbase交易
1. LockOutput: coinbase交易的TXOutput.ScriptPubKey = to.hashPubKey


# UTXO交易 from，to
1. GetWallets
2. GetWallet(from) 得到from的钱包
3. 从钱包地址Base58Decode出hashPubKey
4. Unlock: TXOutput.ScriptPubKey == from.hashPubKey
5. 根据hashPubKey找到未花费的交易输出
6. TXInput.ScriptSig = from.PubKeyBytes（判断交易的输出是否消费)
7. TXOutput.ScriptPubKey = to.hashPubKey

总结：
1. TXInput.ScriptSig 存储钱包的PubKeyBytes公钥，why?公钥验证
2. TXOutput.ScriptPubKey 存储地址的hashPubKey，address经过Base58Decode得到的，如果只存储公钥，那么只要知道公钥的人，就可以改变交易的发起者from


# test
1. ./main createwallets
2. ./main addwallet
   wallet address:  16wyiD52YPBTkodhub1V1fWbc4D8nAqnbd
   wallet address:  13qr7AuvNCuNKA5R7QkwXrG1Ewoy9WSWgw
   wallet address:  1FSj1yEinJML3YvDh4EVUkxX71LYDEbZR2

3. ./main createblockchain -address 16wyiD52YPBTkodhub1V1fWbc4D8nAqnbd
4. ./main getbalance -address 16wyiD52YPBTkodhub1V1fWbc4D8nAqnbd
   Balance of '16wyiD52YPBTkodhub1V1fWbc4D8nAqnbd': 10

5. ./main send -from 16wyiD52YPBTkodhub1V1fWbc4D8nAqnbd -to 13qr7AuvNCuNKA5R7QkwXrG1Ewoy9WSWgw -amount 4

6. ./main send -from 16wyiD52YPBTkodhub1V1fWbc4D8nAqnbd -to 1FSj1yEinJML3YvDh4EVUkxX71LYDEbZR2 -amount 3
