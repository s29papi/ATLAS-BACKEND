const API_KEY = '72a5b4b0-e727-48be-8aa1-5da9d62fe635'; // SOCKET PUBLIC API KEY

const polygonIdx = 4
const gnosisIdx = 3


async function polygonId() {
    const supportedChainsAndRefuellingLimits = await getSupportedChainsAndRefuellingLimits();
    return supportedChainsAndRefuellingLimits.result[polygonIdx].chainId
}
async function gnosisId() {
    const supportedChainsAndRefuellingLimits = await getSupportedChainsAndRefuellingLimits();
    return supportedChainsAndRefuellingLimits.result[gnosisIdx].chainId
}

async function polygonGnosisMinAmount() {
    const supportedChainsAndRefuellingLimits = await getSupportedChainsAndRefuellingLimits();
    return supportedChainsAndRefuellingLimits.result[polygonIdx].limits[2].minAmount
}

async function polygonGnosisMaxAmount() {
    const supportedChainsAndRefuellingLimits = await getSupportedChainsAndRefuellingLimits();
    return supportedChainsAndRefuellingLimits.result[polygonIdx].limits[2].maxAmount
}

async function getQuote() {
    let polygonChainId = await polygonId()
    let gnosisChainId = await gnosisId()
    let sourceChain = polygonChainId;
    let destinationChain = gnosisChainId;
    let fromAssetAddress = '0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee';
    let toAssetAddress = '0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee';
    const amount = 0; // 0.1 Matic decimals
    // const amount = 1 00 000 000 000 000 000; // 0.1 Matic decimals
                    //  21 874 915 496 893 470
                    //  47 382 823 236 483 390
                    //  38 230 216 531 136 629
                //  34 864 248 437 678 900 000
    //                 20211737900598360 
    let recipientAddr = "0x5E0f293eEBa536e6cDeB0B9da03d1b5335dC29De";
    const uniqueRoutesPerBridge = true;
    const sort = "output";
    const singleTxOnly = true;
    const bridgeWithGas = true;

    const qouteUrl = `https://api.socket.tech/v2/quote?bridgeWithGas=${bridgeWithGas}&fromChainId=${sourceChain}&fromTokenAddress=${fromAssetAddress}&toChainId=${destinationChain}&toTokenAddress=${toAssetAddress}&fromAmount=${amount}&userAddress=${recipientAddr}&uniqueRoutesPerBridge=${uniqueRoutesPerBridge}&sort=${sort}&singleTxOnly=${singleTxOnly}`;
    const response = await fetch(qouteUrl, {
        method: `GET`,
        headers: {
            'API-KEY': API_KEY,
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        }
    })
    const json = await response.json();
    return json;
}


async function getSupportedChainsAndRefuellingLimits() {
    const response = await fetch(`https://refuel.socket.tech/chains`, {
        method: `GET`,
        headers: {
            'API-KEY': API_KEY,
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        }
    })
    const json = await response.json();
    return json;
}

// let polygonChainId = await polygonId()
// let gnosisChainId = await gnosisId()
// let sourceChain = polygonChainId;
// let destinationChain = gnosisChainId;

// let fromAssetAddress = '0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee';
// let toAssetAddress = '0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee';
// const amount = 100000000000000000; // 0.1 Matic decimals
// let recipientAddr = "0x5E0f293eEBa536e6cDeB0B9da03d1b5335dC29De";
// const uniqueRoutesPerBridge = true;
// const sort = "output";
// const singleTxOnly = true;
// const bridgeWithGas = true;

async function main() {
    let quote = await getQuote()


    // let supportedChainsAndRefuellingLimits = await getSupportedChainsAndRefuellingLimits()
    // console.log(supportedChainsAndRefuellingLimits.result[4])
    console.log(quote.result)
 

    let maxamount =  await polygonGnosisMaxAmount()

    console.log(maxamount)
    //     10_550_677_389_327_405 16 decimal 
    // 28_730_061_958_472_300_000

    // console.log(destinationChain)
    // console.log(supportedChainsAndRefuellingLimits.result[polygonIdx].limits[2].maxAmount)

    //  '5_139_577_172_852_925',
    //  '17_131_923_909_509_750_000'

}

main();














// const ethers = require("ethers");

// const API_KEY = '72a5b4b0-e727-48be-8aa1-5da9d62fe635'; // SOCKET PUBLIC API KEY

// // Makes a GET request to Socket APIs for quote
// async function getQuote(fromChainId, fromTokenAddress, toChainId, toTokenAddress, fromAmount, userAddress, uniqueRoutesPerBridge, sort, singleTxOnly) {
//     const response = await fetch(`https://api.socket.tech/v2/quote?fromChainId=${fromChainId}&fromTokenAddress=${fromTokenAddress}&toChainId=${toChainId}&toTokenAddress=${toTokenAddress}&fromAmount=${fromAmount}&userAddress=${userAddress}&uniqueRoutesPerBridge=${uniqueRoutesPerBridge}&sort=${sort}&singleTxOnly=${singleTxOnly}`, {
//         method: 'GET',
//         headers: {
//             'API-KEY': API_KEY,
//             'Accept': 'application/json',
//             'Content-Type': 'application/json'
//         }
//     });

//     const json = await response.json();
//     return json;
// }

// // Makes a POST request to Socket APIs for swap/bridge transaction data
// async function getRouteTransactionData(route) {
//     const response = await fetch('https://api.socket.tech/v2/build-tx', {
//         method: 'POST',
//         headers: {
//             'API-KEY': API_KEY,
//             'Accept': 'application/json',
//             'Content-Type': 'application/json'
//         },
//         body: JSON.stringify({ "route": route })
//     });

//     const json = await response.json();
//     return json;
// }

// // GET request to check token allowance given to allowanceTarget by owner
// async function checkAllowance(chainId, owner, allowanceTarget, tokenAddress) {
//     const response = await fetch(`https://api.socket.tech/v2/approval/check-allowance?chainID=${chainId}&owner=${owner}&allowanceTarget=${allowanceTarget}&tokenAddress=${tokenAddress}`, {
//         method: 'GET',
//         headers: {
//             'API-KEY': API_KEY,
//             'Accept': 'application/json',
//             'Content-Type': 'application/json'
//         }
//     });

//     const json = await response.json();
//     return json;
// }

// // Fetches transaction data for token approval 
// async function getApprovalTransactionData(chainId, owner, allowanceTarget, tokenAddress, amount) {
//     const response = await fetch(`https://api.socket.tech/v2/approval/build-tx?chainID=${chainId}&owner=${owner}&allowanceTarget=${allowanceTarget}&tokenAddress=${tokenAddress}&amount=${amount}`, {
//         method: 'GET',
//         headers: {
//             'API-KEY': API_KEY,
//             'Accept': 'application/json',
//             'Content-Type': 'application/json'
//         }
//     });

//     const json = await response.json();
//     return json;
// }

// // Fetches status of the bridging transaction
// async function getBridgeStatus(transactionHash, fromChainId, toChainId) {
//     const response = await fetch(`https://api.socket.tech/v2/bridge-status?transactionHash=${transactionHash}&fromChainId=${fromChainId}&toChainId=${toChainId}`, {
//         method: 'GET',
//         headers: {
//             'API-KEY': API_KEY,
//             'Accept': 'application/json',
//             'Content-Type': 'application/json'
//         }
//     });

//     const json = await response.json();
//     return json;
// }

// // Main function 
// async function main() {

//     // Uses web3 wallet in browser as provider
//     const provider = new ethers.providers.Web3Provider(window.ethereum, "any");

//     // Prompt user for account connections
//     await provider.send("eth_requestAccounts", []);

//     // Stores signer
//     const signer = provider.getSigner();

//     // Bridging Params fetched from users
//     const fromChainId = 137;
//     const toChainId = 56;
//     const fromAssetAddress = "0x2791bca1f2de4661ed88a30c99a7a9449aa84174";
//     const toAssetAddress = "0x55d398326f99059fF775485246999027B3197955";
//     const amount = 100000000; // 100 USDC, USDC is 6 decimals
//     const userAddress = "0x3e8cB4bd04d81498aB4b94a392c334F5328b237b";
//     const uniqueRoutesPerBridge = true; // Returns the best route for a given DEX / bridge combination
//     const sort = "output"; // "output" | "gas" | "time"
//     const singleTxOnly = true;

//     // Quote for bridging 100 USDC on Polygon to USDT on BSC
//     // For single transaction bridging, mark singleTxOnly flag as true in query params
//     const quote = await getQuote(fromChainId,
//         fromAssetAddress, toChainId,
//         toAssetAddress, amount,
//         userAddress, uniqueRoutesPerBridge, sort, singleTxOnly
//     );

//     // Choosing first route from the returned route results 
//     const route = quote.result.routes[0];

//     // Fetching transaction data for swap/bridge tx
//     const apiReturnData = await getRouteTransactionData(route);

//     // Used to check for ERC-20 approvals
//     const approvalData = apiReturnData.result.approvalData;
//     const { allowanceTarget, minimumApprovalAmount } =approvalData;

//     // approvalData from apiReturnData is null for native tokens 
//     // Values are returned for ERC20 tokens but token allowance needs to be checked
//     if (approvalData !== null) {
//         // Fetches token allowance given to Socket contracts
//         const allowanceCheckStatus = await checkAllowance(fromChainId, userAddress, allowanceTarget, fromAssetAddress)
//         const allowanceValue = allowanceCheckStatus.result?.value;

//         // If Socket contracts don't have sufficient allowance
//         if (minimumApprovalAmount > allowanceValue) {
//             // Approval tx data fetched
//             const approvalTransactionData = await getApprovalTransactionData(fromChainId, userAddress, allowanceTarget, fromAssetAddress, amount);

//             const gasPrice = await signer.getGasPrice();

//             const gasEstimate = await provider.estimateGas({
//                 from: signer.address,
//                 to: approvalTransactionData.result?.to,
//                 value: '0x00',
//                 data: approvalTransactionData.result?.data,
//                 gasPrice: gasPrice
//             });

//             const tx = await signer.sendTransaction({
//                 from: approvalTransactionData.result?.from,
//                 to: approvalTransactionData.result?.to,
//                 value: '0x00',
//                 data: approvalTransactionData.result?.data,
//                 gasPrice: gasPrice,
//                 gasLimit: gasEstimate
//             });

//             // Initiates approval transaction on user's frontend which user has to sign
//             const receipt = await tx.wait();

//             console.log('Approval Transaction Hash :', receipt.transactionHash);
//         }
//     }

//     const gasPrice = await signer.getGasPrice();

//     const gasEstimate = await provider.estimateGas({
//         from: signer.address,
//         to: apiReturnData.result.txTarget,
//         value: apiReturnData.result.value,
//         data: apiReturnData.result.txData,
//         gasPrice: gasPrice
//     });

//     const tx = await signer.sendTransaction({
//         from: signer.address,
//         to: apiReturnData.result.txTarget,
//         data: apiReturnData.result.txData,
//         value: apiReturnData.result.value,
//         gasPrice: gasPrice,
//         gasLimit: gasEstimate
//     });

//     // Initiates swap/bridge transaction on user's frontend which user has to sign
//     const receipt = await tx.wait();

//     const txHash = receipt.transactionHash;

//     console.log('Bridging Transaction : ', receipt.transactionHash);

//     // Checks status of transaction every 20 secs
//     const txStatus = setInterval(async () => {
//         const status = await getBridgeStatus(txHash, fromChainId, toChainId);

//         console.log(`SOURCE TX : ${status.result.sourceTxStatus}\nDEST TX : ${status.result.destinationTxStatus}`)

//         if (status.result.destinationTxStatus == "COMPLETED") {
//             console.log('DEST TX HASH :', status.result.destinationTransactionHash);
//             clearInterval(txStatus);
//         }
//     }, 20000);

// }

// main();