import { FrameRequest, getFrameMessage } from '@coinbase/onchainkit';
import { NextRequest, NextResponse } from 'next/server';

async function getResponse(req: NextRequest): Promise<NextResponse> {
  const body: FrameRequest = await req.json();
  const searchParams = req.nextUrl.searchParams;
  const gameId:any = searchParams.get("gameId");
  const gameName:any = searchParams.get("gameName");
  const gameSetup:any = searchParams.get("gameSetup");
  const stakeAmount:any = searchParams.get("stakeAmount");
  const creatorFid:any = searchParams.get("creatorFid");
  const buttonId = body.untrustedData.buttonIndex;
  
  let queryParams = `gameId=${gameId}&&gameName=${gameName}&&gameSetup=${gameSetup}&&stakeAmount=${stakeAmount}&&creatorFid=${creatorFid}`
  let button2ImageUrl = "https://wag3r-bot-gamma.vercel.app/og/account?" + `${queryParams}`
  let button2PostUrl = "https://wag3r-bot-gamma.vercel.app/api/frame/account?" + `${queryParams}`
  if (buttonId == 1) {
    return new NextResponse(`<!DOCTYPE html><html><head>
    <title>Start My Match</title>
    <meta property="fc:frame" content="vNext" />
    <meta property="fc:frame:image" content="https://wag3r-bot.vercel.app/og"/>
    <meta property="fc:frame:button:1" content="Back" />
    <meta property="fc:frame:button:1:action" content="post"/>
    <meta property="fc:frame:button:2" content="Stake" />
    <meta property="fc:frame:button:2:action" content="post"/>
    <meta property="fc:frame:post_url" content="https://wag3r-bot.vercel.app/api/frame/stake"/>
</head></html>`);
  }

  // if (buttonId == 2) {
  //   return new NextResponse(`<!DOCTYPE html><html><head>
  //        <title>Account</title>
  //        <meta property="fc:frame" content="vNext" />
  //        <meta property="fc:frame:image" content="${button2ImageUrl}"/>
  //        <meta property="fc:frame:button:1" content="Back" />
  //        <meta property="fc:frame:button:1:action" content="post"/>
  //        <meta property="fc:frame:button:2" content="Withdraw" />
  //        <meta property="fc:frame:button:2:action" content="post_redirect"/>
  //        <meta property="fc:frame:button:3" content="Deposit" />
  //        <meta property="fc:frame:button:3:action" content="post_redirect"/>
  //        <meta property="fc:frame:button:4" content="Refresh" />
  //        <meta property="fc:frame:button:4:action" content="post"/>
  //        <meta property="fc:frame:post_url" content="${button2PostUrl}"/>
  //       </head></html>`);
  // }





  

  // handles accounts
  return NextResponse.redirect("https://wag3r-bot.vercel.app/");
}

export async function POST(req: NextRequest): Promise<Response> {
  return getResponse(req);
}

export const dynamic = 'force-dynamic';



function ifAccountBalanceIsEqualGreaterStakeAmount() {
    return new NextResponse(`<!DOCTYPE html><html><head>
            <title>Start My Match</title>
            <meta property="fc:frame" content="vNext" />
            <meta property="fc:frame:image" content="https://wag3r-bot.vercel.app/og"/>
            <meta property="fc:frame:button:1" content="Back" />
            <meta property="fc:frame:button:1:action" content="post"/>
            <meta property="fc:frame:button:2" content="Stake" />
            <meta property="fc:frame:button:2:action" content="post"/>
            <meta property="fc:frame:post_url" content="https://wag3r-bot.vercel.app/api/frame/stake"/>
        </head></html>`);
}




//   // Step 3. Validate the message
//   const { isValid, message } = await getFrameMessage(body, {neynarApiKey: "NEYNAR_ONCHAIN_KIT"});
//   let following;
//   let liked;
//   if (message?.following) following = message?.following
//   if (message?.liked) liked = message?.liked

//   // if the interactive user hasn't followed and liked cast return 
//   if (!liked && !following) {
//           return new NextResponse(`<!DOCTYPE html><html><head>
//           <title>Like & Follow</title>
//           <meta property="fc:frame" content="vNext" />
//           <meta property="fc:frame:image" content="https://frames-follow-like-refuel.vercel.app/second-page.png"/>
//           <meta property="fc:frame:button:1" content="Refuel" />
//           <meta property="fc:frame:button:1:action" content="post"/>
//           <meta property="fc:frame:post_url" content="https://frames-follow-like-refuel.vercel.app/api/frame"/>
//       </head></html>`);
//   }
// // if the interactive user hasn't liked cast return 
//   if (!liked && following) {
//         return new NextResponse(`<!DOCTYPE html><html><head>
//         <title>Like Cast</title>
//         <meta property="fc:frame" content="vNext" />
//         <meta property="fc:frame:image" content="https://frames-follow-like-refuel.vercel.app/fourth-page.png"/>
//         <meta property="fc:frame:button:1" content="Refuel" />
//         <meta property="fc:frame:button:1:action" content="post"/>
//         <meta property="fc:frame:post_url" content="https://frames-follow-like-refuel.vercel.app/api/frame"/>
//     </head></html>`);
//   }

//   if (liked && !following) {
//         return new NextResponse(`<!DOCTYPE html><html><head>
//         <title>Follow User</title>
//         <meta property="fc:frame" content="vNext" />
//         <meta property="fc:frame:image" content="https://frames-follow-like-refuel.vercel.app/third-page.png"/>
//         <meta property="fc:frame:button:1" content="Refuel" />
//         <meta property="fc:frame:button:1:action" content="post"/>
//         <meta property="fc:frame:post_url" content="https://frames-follow-like-refuel.vercel.app/api/frame"/>
//     </head></html>`);
//   }
//   // in sedn frames, process refuel and return successful 
//   //   return new NextResponse(`<!DOCTYPE html><html><head>
//   //             <title>Input Wallet Address</title>
//   //             <meta property="fc:frame" content="vNext" />
//   //             <meta property="fc:frame:image" content="https://magenta-hollow-tiglon-795.mypinata.cloud/ipfs/QmVfJeE5pEXPhALNSj7a7a2EJ9MDEG7MnhegQbBhFeMEVc"/>
//   //             <meta property="fc:frame:input:text" content="Wallet Address..." />
//   //             <meta property="fc:frame:button:1" content="Submit" />
//   //             <meta property="fc:frame:button:1:action" content="post"/>
//   //             <meta property="fc:frame:post_url" content="https://frames-follow-like-refuel.vercel.app/api/frame/refuel"/>
//   // </head></html>`);

