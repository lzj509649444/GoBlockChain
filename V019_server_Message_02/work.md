https://en.bitcoin.it/wiki/Protocol_documentation#inv

Messages: version,verack,addr

1. block Height,MineBlock


server_message_02.go

NODE_ID=4000 ./main createwallets

NODE_ID=4000 ./main addwallet
wallet address:  1EVT8d6YBvfeE4SYCCRfoWL6ZXk9A3hPU1

NODE_ID=4000 ./main addwallet
wallet address:  1Ndq99Y81UuVJ11jwMrMNN17kXMMMYh8Ez

NODE_ID=4000 ./main addwallet
wallet address:  17pNW3EuirpDDXyrY5oqBCuMvRNL27uKjk


NODE_ID=4000 ./main createblockchain -address 1EVT8d6YBvfeE4SYCCRfoWL6ZXk9A3hPU1

NODE_ID=4000 ./main getbalance -address 1EVT8d6YBvfeE4SYCCRfoWL6ZXk9A3hPU1

NODE_ID=4000 ./main send -from 1EVT8d6YBvfeE4SYCCRfoWL6ZXk9A3hPU1 -to 1Ndq99Y81UuVJ11jwMrMNN17kXMMMYh8Ez -amount 3

NODE_ID=4000 ./main getbalance -address 1Ndq99Y81UuVJ11jwMrMNN17kXMMMYh8Ez
