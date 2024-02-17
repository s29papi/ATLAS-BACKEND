import { ImageResponse } from 'next/server'
import Image from 'next/image'
import baseStake from '../../../public/Base-Stake-Image-rescale.png'
// App router includes @vercel/og.
// No need to install it.

export const runtime = 'edge';

export async function GET() {
    return new ImageResponse(
      (
        <Image
        src={baseStake}
        alt="Picture of the author"
        sizes="100vw"
        style={{
          width: '100%',
          height: 'auto',
        }}
         />
      )
    )
  }