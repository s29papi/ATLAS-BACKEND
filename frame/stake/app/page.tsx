import { getFrameMetadata } from '@coinbase/onchainkit';
import type { Metadata } from 'next';



const frameMetadata = getFrameMetadata({
  buttons: [
      {label: 'Accept Challenge & Stake Tokens', action: 'post_redirect'},
  ],
  image: 'https://wag3r-bot-stake.vercel.app/stadium-figma-test-1.png',
  post_url: 'https://wag3r-bot-stake.vercel.app/api',
});

export const metadata: Metadata = {
  title: 'Refuel-Frame by socket.',
  description: 'Follow this user, Like the post, and Refuel.',
  openGraph: {
    title: 'Refuel-Frame by socket.',
    description: 'Follow this user, Like the post, and Refuel.',
    images: [`https://wag3r-bot-stake.vercel.app/stadium-figma-test-1.png`],
  },
  other: {
    ...frameMetadata,
  },
};



// export default async function Page(props: Props) {

//   return (
//     <>
//       <h1>Refuel-Frame by socket.</h1>
//     </>
//   );
// }


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

