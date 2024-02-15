import { getFrameMetadata } from '@coinbase/onchainkit';
import type { Metadata } from 'next';



const frameMetadata = getFrameMetadata({
  buttons: [
      {label: 'Wager', action: 'post'},
      {label: 'Un-Wager', action: 'post'},
      {label: 'Info', action: 'post'},
  ],
  image: 'https://wag3r-bot.vercel.app/stadium-first-page.png',
  post_url: 'https://wag3r-bot.vercel.app/api/frame',
});

export const metadata: Metadata = {
  title: 'Refuel-Frame by socket.',
  description: 'Follow this user, Like the post, and Refuel.',
  openGraph: {
    title: 'Refuel-Frame by socket.',
    description: 'Follow this user, Like the post, and Refuel.',
    images: [`https://wag3r-bot.vercel.app/stadium-first-page.png`],
  },
  other: {
    ...frameMetadata,
  },
};

type Props = {
  params: {},
  searchParams: { [key: string]: string}
}

export default async function Page(props: Props) {
  function getSearchParams() {
    const searchParams = props.searchParams;
    console.log(searchParams["game-id"])
    console.log(4949)
  }

  return (
    <>
      <h1>Refuel-Frame by socket.</h1>
      <button onClick={getSearchParams} style={{ borderRadius: '20px', backgroundColor: 'red', color: 'white', padding: '10px 20px', border: 'none' }}>Finish</button>
    </>
  );
}


// ref: https://www.pinata.cloud/blog/how-to-make-a-frame-on-farcaster-using-ipfs
// spec: https://docs.farcaster.xyz/reference/frames/spec
// playground: https://og-playground.vercel.app/

