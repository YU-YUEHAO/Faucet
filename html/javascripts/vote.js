
// const web3 = new Web3(web3.givenProvider); 
let web3 = new Web3(Web3.givenProvider || new Web3.providers.HttpProvider("http://localhost:8545"));
console.log(web3);
let yu=[];

 var accouts=web3.eth.getAccounts().then(function(accounts){
       yu=accounts;
      console.log(yu[0])
 });
const abi=[
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "string",
				"name": "eventType",
				"type": "string"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "_proposalId",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "_limitTime",
				"type": "uint256"
			}
		],
		"name": "ProposeEvt",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "string",
				"name": "eventType",
				"type": "string"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "_voter",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "timestamp",
				"type": "uint256"
			}
		],
		"name": "VoteEvt",
		"type": "event"
	},
	{
		"inputs": [
			{
				"internalType": "string",
				"name": "_pName",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "_pCtx",
				"type": "string"
			},
			{
				"internalType": "uint256",
				"name": "_limitTime",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "pId",
				"type": "uint256"
			}
		],
		"name": "createProposal",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "pId",
				"type": "uint256"
			}
		],
		"name": "doVoting",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "getBlockTime",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "t",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "pId",
				"type": "uint256"
			}
		],
		"name": "getProposalCtx",
		"outputs": [
			{
				"internalType": "string",
				"name": "s",
				"type": "string"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "pId",
				"type": "uint256"
			}
		],
		"name": "getProposalLimit",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "t",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "pId",
				"type": "uint256"
			}
		],
		"name": "getProposalName",
		"outputs": [
			{
				"internalType": "string",
				"name": "s",
				"type": "string"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "pId",
				"type": "uint256"
			}
		],
		"name": "getProposalVCnt",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "v",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"name": "proposals",
		"outputs": [
			{
				"internalType": "string",
				"name": "pName",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "pCtx",
				"type": "string"
			},
			{
				"internalType": "address",
				"name": "chairperson",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "voteCount",
				"type": "uint256"
			},
			{
				"internalType": "bool",
				"name": "initialized",
				"type": "bool"
			},
			{
				"internalType": "uint256",
				"name": "limitTime",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "pId",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "pId",
				"type": "uint256"
			},
			{
				"internalType": "address",
				"name": "voterAddr",
				"type": "address"
			}
		],
		"name": "queryVoting",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	}
]
/*
$("#btn_confirmTransaction").click(function(){
    txid_=$("#txid").val();
     console.log(txid_);
    var myContract = new web3.eth.Contract(abi, '0xb253E02584402a8f226c1768Fd1A2D051D7F75aa');
    myContract.methods.confirmTransaction(txid_).send({from:yu[0],gas:1500000,gasPrice:'30000000'})
    .on('transactionHash',function(receipt){
        console.log("transactionHash："+receipt)
    })
    .on('receipt',function(receipt){
        console.log(receipt)
    })
    .on('confirmation',function(confirmation){
        console.log(confirmation)
    })
    .on('error',function(error){
        console.log(error)
    })
})
*/