package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// CLI responsible for processing command line arguments
type CLI struct{}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  createblockchain -address ADDRESS - Create a blockchain and send genesis block reward to ADDRESS")
	fmt.Println("  printchain - Print all the blocks of the blockchain")
	fmt.Println("  send -from FROM -to TO -amount AMOUNT - Send AMOUNT of coins from FROM address to TO")
	fmt.Println("  getbalance -address ADDRESS - Get balance of ADDRESS")
	// fmt.Println("  createwallet - Create a Wallet")
	// fmt.Println("  walletAddress - Get Wallet Address")
	fmt.Println("  createwallets - Create empty wallets")
	fmt.Println("  addwallet - Add a Wallet To wallets")
	fmt.Println("  printwallets - Print all Wallets Address From Wallets")
	fmt.Println("  reindexutxo - Rebuilds the UTXO set")
	fmt.Println("  startnode -node-id NODE_ID - Start a node with specified ID")

}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

// Run parses command line arguments and processes commands
func (cli *CLI) Run() {
	cli.validateArgs()

	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
	// createWalletCmd := flag.NewFlagSet("createwallet", flag.ExitOnError)
	// walletAddressCmd := flag.NewFlagSet("walletAddress", flag.ExitOnError)
	reindexUTXOCmd := flag.NewFlagSet("reindexutxo", flag.ExitOnError)

	createWalletsCmd := flag.NewFlagSet("createwallets", flag.ExitOnError)
	addWalletCmd := flag.NewFlagSet("addwallet", flag.ExitOnError)
	printWalletsCmd := flag.NewFlagSet("printwallets", flag.ExitOnError)

	startNodeCmd := flag.NewFlagSet("startnode", flag.ExitOnError)

	createBlockchainAddress := createBlockchainCmd.String("address", "", "The address to send genesis block reward to")
	sendFrom := sendCmd.String("from", "", "Source wallet address")
	sendTo := sendCmd.String("to", "", "Destination wallet address")
	sendAmount := sendCmd.Int("amount", 0, "Amount to send")
	getBalanceAddress := getBalanceCmd.String("address", "", "The address to get balance for")
	startNodeID := startNodeCmd.Int("node-id", 0, "Node ID")

	switch os.Args[1] {
	case "createblockchain":
		err := createBlockchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "send":
		err := sendCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "getbalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	// case "createwallet":
	// 	err := createWalletCmd.Parse(os.Args[2:])
	// 	if err != nil {
	// 		log.Panic(err)
	// 	}
	// case "walletAddress":
	// 	err := walletAddressCmd.Parse(os.Args[2:])
	// 	if err != nil {
	// 		log.Panic(err)
	// 	}
	case "createwallets":
		err := createWalletsCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "addwallet":
		err := addWalletCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printwallets":
		err := printWalletsCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "reindexutxo":
		err := reindexUTXOCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "startnode":
		err := startNodeCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if createBlockchainCmd.Parsed() {
		if *createBlockchainAddress == "" {
			createBlockchainCmd.Usage()
			os.Exit(1)
		}
		cli.createBlockchain(*createBlockchainAddress)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}

	if sendCmd.Parsed() {
		if *sendFrom == "" || *sendTo == "" || *sendAmount <= 0 {
			sendCmd.Usage()
			os.Exit(1)
		}

		cli.send(*sendFrom, *sendTo, *sendAmount)
	}

	if getBalanceCmd.Parsed() {
		if *getBalanceAddress == "" {
			getBalanceCmd.Usage()
			os.Exit(1)
		}
		cli.getBalance(*getBalanceAddress)
	}

	// if createWalletCmd.Parsed() {
	// 	cli.createWallet()
	// }

	// if walletAddressCmd.Parsed() {
	// 	cli.walletAddress()
	// }

	if createWalletsCmd.Parsed() {
		cli.createWallets()
	}

	if addWalletCmd.Parsed() {
		cli.addWallet()
	}

	if printWalletsCmd.Parsed() {
		cli.printWallets()
	}

	if reindexUTXOCmd.Parsed() {
		cli.reindexUTXO()
	}

	if startNodeCmd.Parsed() {
		if *startNodeID == 0 {
			startNodeCmd.Usage()
			os.Exit(1)
		}
		cli.startNode(*startNodeID)
	}

}
