import { NextRequest, NextResponse } from 'next/server';


async function getResponse(req: NextRequest): Promise<NextResponse> {
  const data = await req.json(); 
  const searchParams = req.nextUrl.searchParams
  const gameid:any = searchParams.get("gameid")
  return NextResponse.redirect('https://wag3r-bot-stake.vercel.app/~/stake?gameid=' + `${gameid}`, {status: 302});
}

export async function POST(req: NextRequest): Promise<Response> {
  return getResponse(req);
}

export const dynamic = 'force-dynamic';