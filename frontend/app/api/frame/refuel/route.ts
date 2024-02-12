import { NextRequest, NextResponse } from 'next/server';
import { ethers } from 'ethers';
import { Alchemy, Network } from 'alchemy-sdk';


async function getResponse(req: NextRequest): Promise<NextResponse> {
  // when a user swap+bridges to get gas tokens on the destination chain 
  // 1. the user is going to recieve his token by / through the bridge and das 

  // process refuel with socket api's
  return new NextResponse(`<!DOCTYPE html><html><head>
  <title>Success</title>
  <meta property="fc:frame" content="vNext" />
  <meta property="fc:frame:image" content="https://magenta-hollow-tiglon-795.mypinata.cloud/ipfs/QmXGeear4FPk4YAjU3hGTFZCSXi8dWV9oF9njxnyXCsGb5"/>
  <meta property="fc:frame:button:1" content="Success 🎉" />
</head></html>`);
  // const data = await req.json();
  // let getParams = req.nextUrl.searchParams
  
  // const searchParams = req.nextUrl.searchParams
  // const id:any = searchParams.get("id")
  // const idAsNumber = parseInt(id)

  // if (idAsNumber === 0){
  //   const destinationAddress = data.untrustedData.inputText;
  //   let userInfo = "DestinationAddress=" + destinationAddress; 
  //   let response = new NextResponse(`<!DOCTYPE html><html><head>
  //       <title>Frame Sends Eth</title>
  //       <meta property="fc:frame" content="vNext" />
  //       <meta property="fc:frame:image" content="https://magenta-hollow-tiglon-795.mypinata.cloud/ipfs/QmZPrZ45GrnmjbGw6Xj27mzgpju7FCguKAbwBkUVxBTPVB"/>
  //       <meta property="fc:frame:button:1" content="Select Destination Chain" />
  //       <meta property="fc:frame:button:1:action" content="post"/>
  //       <meta property="fc:frame:post_url" content="https://socket-pay.vercel.app/api/send-frame?id=1&${userInfo}"/>
  //     </head></html>`);
  //   return response
  // }

  // let nextId = idAsNumber + 1

  // if(idAsNumber === 1){
  //   let destinationAddress = searchParams.get("DestinationAddress");
  //   let userInfo = "DestinationAddress=" + destinationAddress; 
  //     return new NextResponse(`<!DOCTYPE html><html><head>
  //         <title>Select the Destination chain...</title>
  //         <meta property="fc:frame" content="vNext" />
  //         <meta property="fc:frame:image" content="https://magenta-hollow-tiglon-795.mypinata.cloud/ipfs/QmZPrZ45GrnmjbGw6Xj27mzgpju7FCguKAbwBkUVxBTPVB"/>
  //         <meta property="fc:frame:button:1" content="Ethereum" />
  //         <meta property="fc:frame:button:1:action" content="post" />
  //         <meta property="fc:frame:button:2" content="Arbitrum" />
  //         <meta property="fc:frame:button:2:action" content="post" />
  //         <meta property="fc:frame:button:3" content="Optimism" />
  //         <meta property="fc:frame:button:3:action" content="post" />
  //         <meta property="fc:frame:post_url" content="https://socket-pay.vercel.app/api/send-frame?id=${nextId}&${userInfo}" />
  //       </head></html>`);
  // } 

  // nextId = idAsNumber + 1

  // if(idAsNumber === 2){
  //   let destinationAddress = searchParams.get("DestinationAddress");
  //   let userInfo = "DestinationAddress=" + destinationAddress + "&" + "DestinationChainId="; 
  //   let buttonId = data.untrustedData.buttonIndex;
  //   if (buttonId === 1){ 
  //     userInfo += "1" 
  //   }
  //   if (buttonId === 2){ 
  //     userInfo += "42161"
  //   }
  //   if (buttonId === 3){ 
  //     userInfo += "10" 
  //   }
  //   return new NextResponse(`<!DOCTYPE html><html><head>
  //       <title>Frame Sends Eth</title>
  //       <meta property="fc:frame" content="vNext" />
  //       <meta property="fc:frame:image" content="https://magenta-hollow-tiglon-795.mypinata.cloud/ipfs/QmZPrZ45GrnmjbGw6Xj27mzgpju7FCguKAbwBkUVxBTPVB"/>
  //       <meta property="fc:frame:button:1" content="Select Source Chain" />
  //       <meta property="fc:frame:button:1:action" content="post"/>
  //       <meta property="fc:frame:post_url" content="https://socket-pay.vercel.app/api/send-frame?id=${nextId}&${userInfo}"/>
  //     </head></html>`);
  //   } 

  //   nextId = idAsNumber + 1

  //   if(idAsNumber === 3){
  //     let destinationAddress = searchParams.get("DestinationAddress");
  //     let chainId = searchParams.get("DestinationChainId");
  //     let userInfo = "DestinationAddress=" + destinationAddress + "&" + "DestinationChainId=" + chainId; 

  //     return new NextResponse(`<!DOCTYPE html><html><head>
  //       <title>Select the Source chain...</title>
  //       <meta property="fc:frame" content="vNext" />
  //       <meta property="fc:frame:image" content="https://magenta-hollow-tiglon-795.mypinata.cloud/ipfs/QmZPrZ45GrnmjbGw6Xj27mzgpju7FCguKAbwBkUVxBTPVB"/>
  //       <meta property="fc:frame:button:1" content="Ethereum" />
  //       <meta property="fc:frame:button:1:action" content="post" />
  //       <meta property="fc:frame:button:2" content="Arbitrum" />
  //       <meta property="fc:frame:button:2:action" content="post" />
  //       <meta property="fc:frame:button:3" content="Optimism" />
  //       <meta property="fc:frame:button:3:action" content="post" />
  //       <meta property="fc:frame:post_url" content="https://socket-pay.vercel.app/api/send-frame?id=${nextId}&${userInfo}" />
  //     </head></html>`);
  //     } 

  //     nextId = idAsNumber + 1

  //     if(idAsNumber === 4){
  //       let destinationAddress = searchParams.get("DestinationAddress");
  //       let chainId = searchParams.get("DestinationChainId");
  //       let userInfo = "DestinationAddress=" + destinationAddress + "&" + "DestinationChainId=" +  chainId + "&" + "SourceChainId="; 
  //       const buttonId = data.untrustedData.buttonIndex;
  //       if (buttonId === 1){ 
  //         userInfo += "1" 
  //       }
  //       if (buttonId === 2){ 
  //         userInfo += "42161" 
  //       }
  //       if (buttonId === 3){ 
  //         userInfo += "10" 
  //       }
  //       return new NextResponse(`<!DOCTYPE html><html><head>
  //       <title>Frame returns user Eth balance</title>
  //       <meta property="fc:frame" content="vNext" />
  //       <meta property="fc:frame:image" content="https://magenta-hollow-tiglon-795.mypinata.cloud/ipfs/QmZPrZ45GrnmjbGw6Xj27mzgpju7FCguKAbwBkUVxBTPVB"/>
  //       <meta property="fc:frame:input:text" content="Enter Eth Amount" />
  //       <meta property="fc:frame:button:1" content="Proceed to sign Tx" />
  //       <meta property="fc:frame:button:1:action" content="post_redirect" />
  //       <meta property="fc:frame:post_url" content="https://socket-pay.vercel.app/api/send-frame?id=${nextId}&${userInfo}" />
  //     </head></html>`);
  //   }

  //   // populate it to have the input amount
  //   if(idAsNumber === 5){
  //     let destinationAddress = searchParams.get("DestinationAddress");
  //       let dchainId = searchParams.get("DestinationChainId");
  //       let schainId = searchParams.get("SourceChainId");
  //       const inputAmount = data.untrustedData.inputText;
  //       let userInfo = "DestinationAddress=" + destinationAddress + "&" + "DestinationChainId=" +  dchainId + "&" + "SourceChainId=" + schainId + "&" + "InputAmount=" + inputAmount; 
  //     const buttonId = data.untrustedData.buttonIndex;
  //     if (buttonId === 1){ 
  //       let processPlusSign = "redirectPage"
  //       return NextResponse.redirect("https://socket-pay.vercel.app/redirect/" + `${processPlusSign}` + "?" + `${userInfo}` , {status: 302});
  //     }
  //   }
  
  //     return NextResponse.redirect('https://socket-pay.vercel.app/', {status: 302});
}

export async function POST(req: NextRequest): Promise<Response> {
  return getResponse(req);
}

export const dynamic = 'force-dynamic';





// const alchemy = new Alchemy(settings);
// alchemy.core.getBlockNumber().then(console.log);
// //   // MetaMask requires requesting permission to connect users accounts
// // await provider.send("eth_requestAccounts", []);

// // const signer = provider.getSigner()
// const settings = {
  //     apiKey: 'PcjF8aGR1lwutd6YiawYDs05rSYnSL-A', // Replace with your Alchemy API Key.
  //     network: Network.OPT_MAINNET, // Replace with your network.
  // };
  