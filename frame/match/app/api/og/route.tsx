import { ImageResponse } from 'next/server'
// App router includes @vercel/og.
// No need to install it.
import stakeBg from '../../../public/stake.svg'
export const runtime = 'edge';

export async function GET() {
    return new ImageResponse(
      (
        <div
          style={{
            width: '100%',
            height: '100%',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            fontSize: 128,
            background: 'lavender',
            backgroundImage: `url(${stakeBg.src}),`
          }}
        >
          Hello!
        </div>
      )
    )
  }