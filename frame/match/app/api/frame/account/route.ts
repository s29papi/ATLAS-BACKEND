import { FrameRequest, getFrameMessage } from '@coinbase/onchainkit';
import { NextRequest, NextResponse } from 'next/server';

async function getResponse(req: NextRequest): Promise<NextResponse> {
  const body: FrameRequest = await req.json();
  const searchParams = req.nextUrl.searchParams;
  const gameId:any = searchParams.get("gameId");
  const buttonId = body.untrustedData.buttonIndex;
  // pass in facaster Id + hash of transaction generated from the tx
  // bot would resolve the balance
  const { isValid, message } = await getFrameMessage(body , {
    neynarApiKey: 'NEYNAR_ONCHAIN_KIT'
  });

  if (buttonId == 2) {
    return NextResponse.redirect('https://wag3r-bot.vercel.app/~/stake?${message.fid}', {status: 302});
  }

  if (buttonId == 3) {
    return NextResponse.redirect('https://wag3r-bot.vercel.app/~/unstake?${message.fid}', {status: 302});
  }

  // handles the first button which is back
  // and any other externalities
  return NextResponse.redirect("https://wag3r-bot.vercel.app/?gameId=${gameId}");
}

export async function POST(req: NextRequest): Promise<Response> {
  return getResponse(req);
}

export const dynamic = 'force-dynamic';

