import { ImageResponse } from 'next/server'
import Image from 'next/image'
import stakeBG from '../../../public/stake.svg'
// App router includes @vercel/og.
// No need to install it.

export const runtime = 'edge';

export async function GET() {
    return new ImageResponse(
      (
            <Image
            alt="Mountains"
            src={stakeBG}
            placeholder="blur"
            quality={100}
            fill
            sizes="100vw"
            style={{
            objectFit: 'cover',
            }}
        />
      )
    )
  }