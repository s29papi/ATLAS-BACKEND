import { getFrameMetadata } from '@coinbase/onchainkit';
import type { Metadata, ResolvingMetadata } from 'next'

type Props = {
  params: { gameId: string }
  searchParams: { [key: string]: string | string[] | undefined }
}
 

export async function generateMetadata(
  { params, searchParams }: Props,
  parent: ResolvingMetadata
): Promise<Metadata> { 
  const gameid = searchParams["gameId"];
  const gameName = searchParams["gameName"]
  const gameSetup = searchParams["gameSetup"]
  const stakeAmount = searchParams["stakeAmount"]
  const creatorFid = searchParams["creatorFid"]

  let queryParams = `gameId=${gameid}&&gameName=${gameName}&&gameSetup=${gameSetup}&&stakeAmount=${stakeAmount}&&creatorFid=${creatorFid}`
  let postUrl = "https://wag3r-bot.vercel.app/api/frame?" + `${queryParams}`;
  let imageUrl = "https://wag3r-bot-gamma.vercel.app/og/landingframe?" + `${queryParams}`;

  const frameMetadata = getFrameMetadata({
    buttons: [
        {label: 'Accept Challenge', action: 'post'},
        {label: 'Account', action: 'post'},
    ],
    image: imageUrl,
    post_url: postUrl,
  });

  return {
    title: 'Match By Versus.',
    description: 'Frontend Match Management for Versus App.',
    openGraph: {
      title: 'Match By Versus.',
      description: 'Frontend Match Management for Versus App.',
      images: [imageUrl],
    },
    other: {
      ...frameMetadata,
    }
  }
}

// A-New-Challenger-Has-Entered-The-Ring-Resize


export default function Page({
  params,
  searchParams,
}: {
  params: { slug: string };
  searchParams?: { [key: string]: string | string[] | undefined };
}) {
  return <h1>Match By Versus iiiooo.</h1>;
}



// spec: https://docs.farcaster.xyz/reference/frames/spec  
// playground: https://og-playground.vercel.app/


