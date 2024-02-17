import { FrameRequest, getFrameMessage } from '@coinbase/onchainkit';
import { NextRequest, NextResponse } from 'next/server';

async function getResponse(req: NextRequest): Promise<NextResponse> {
  const body: FrameRequest = await req.json();

  
  const searchParams = req.nextUrl.searchParams
  const gameId:any = searchParams.get("game-id")

  const buttonId = body.untrustedData.buttonIndex;
  
  
  if (buttonId == 1) {
    // this is in one of three states

    // get the value from the bot of the user
    // get the value of the game stake amount

    // if it hasn't staked (this values are fetched from golang)

    // if already staked and value is >= game amount

    // if already taked and value is less than 
    return NextResponse.redirect("https://wag3r-bot.vercel.app/")
  }
  // wager: is the second buttonId
  if (buttonId == 2) {
    return new NextResponse(`<!DOCTYPE html><html><head>
          <title>Start My Match</title>
          <meta property="fc:frame" content="vNext" />
          <meta property="fc:frame:image" content="https://wag3r-bot.vercel.app/stadium-figma-test-1.png"/>
          <meta property="fc:frame:button:1" content="Start My Match" />
          <meta property="fc:frame:button:1:action" content="post"/>
          <meta property="fc:frame:post_url" content=""/>
      </head></html>`);
  }

return new NextResponse(`<!DOCTYPE html><html><head>
  <title>Button doesnt exist</title>
</head></html>`);
}

export async function POST(req: NextRequest): Promise<Response> {
  return getResponse(req);
}

export const dynamic = 'force-dynamic';



function ifAccountBalanceIsEqualGreaterStakeAmount() {
    return new NextResponse(`<!DOCTYPE html><html><head>
            <title>Start My Match</title>
            <meta property="fc:frame" content="vNext" />
            <meta property="fc:frame:image" content="https://wag3r-bot.vercel.app/api/og"/>
            <meta property="fc:frame:button:1" content="Back" />
            <meta property="fc:frame:button:1:action" content="post"/>
            <meta property="fc:frame:button:1" content="Stake" />
            <meta property="fc:frame:button:1:action" content="post"/>
            <meta property="fc:frame:post_url" content="https://wag3r-bot.vercel.app/api/frame/stake"/>
        </head></html>`);
}


// <meta property="fc:frame:image" content="https://wag3r-bot.vercel.app/api/og"/>
// <meta property="fc:frame:button:1" content="Back" />
// <meta property="fc:frame:button:1:action" content="post"/>
// <meta property="fc:frame:button:1" content="Stake" />
// <meta property="fc:frame:button:1:action" content="post"/>