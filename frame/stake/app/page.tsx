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
  const gameId = params.gameId;
  let postUrl = "https://wag3r-bot.vercel.app/api?gameId=" + `${gameId}`;

  const frameMetadata = getFrameMetadata({
    buttons: [
        {label: 'Accept Challenge & Stake Tokens', action: 'post'},
    ],
    image: 'https://wag3r-bot-stake.vercel.app/stadium-figma-test-1.png',
    post_url: postUrl,
  });

  return {
    title: 'Stake By Versus.',
    description: 'Frontend Stake Management for Versus App.',
    openGraph: {
      title: 'Stake By Versus.',
      description: 'Frontend Stake Management for Versus App.',
      images: [`https://wag3r-bot-stake.vercel.app/stadium-figma-test-1.png`],
    },
    other: {
      ...frameMetadata,
    }
  }
}



// export const metadata: Metadata = {
//   title: 'Refuel-Frame by socket.',
//   description: 'Follow this user, Like the post, and Refuel.',
//   openGraph: {
//     title: 'Refuel-Frame by socket.',
//     description: 'Follow this user, Like the post, and Refuel.',
//     images: [`https://wag3r-bot-stake.vercel.app/stadium-figma-test-1.png`],
//   },
//   other: {
//     ...frameMetadata,
//   },
// };





export default function Page({
  params,
  searchParams,
}: {
  params: { slug: string };
  searchParams?: { [key: string]: string | string[] | undefined };
}) {
  return <h1>{searchParams?.gameid || "Hello!"}</h1>;
}


// ref: https://www.pinata.cloud/blog/how-to-make-a-frame-on-farcaster-using-ipfs
// spec: https://docs.farcaster.xyz/reference/frames/spec
// playground: https://og-playground.vercel.app/

