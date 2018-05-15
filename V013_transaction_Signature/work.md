1. TXInput
Signature []byte
PubKey    []byte

2. TXOutput
PubKeyHash []byte

3. //tx.SetID()
tx.ID = tx.Hash()

4. blockchain
blockchain.SignTransaction(&tx, wallet.PrivateKey)

5. MineBlock
blockchain.VerifyTransaction(tx)
