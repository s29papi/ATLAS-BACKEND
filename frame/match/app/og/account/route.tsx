/* eslint-disable @next/next/no-img-element  */
/* eslint-disable jsx-ally/alt-text */
// @ts-nocheck
import { ImageResponse } from 'next/server'

export const runtime = 'edge';


export async function GET(req: Request) {
  const imageData = await fetch(new URL('../../../public/base.png', import.meta.url)).then(
    (res) => res.arrayBuffer(),
  );
  const geistData = await fetch(new URL('../../../public/Geist-Bold-BF6569491da5a14.otf', import.meta.url)).then(
    (res) => res.arrayBuffer(),
  );
  return new ImageResponse ((
            <div style={{position: 'relative', display: 'flex'}}>
                    <img 
                        src={imageData}
                    />
                    <div style={{display: 'flex', position: 'absolute', top: '65%', left: '50%', transform: 'translate(-50%, -50%)', textAlign: 'center', color: '#66757F', fontSize: '72px', fontWeight: '1200'}}> 
                            <span style={{fontStyle: 'italic'}}>BALANCE $ 20 USDC</span>
                    </div>
             </div>
        ),
        {
            fonts: [
                {
                    name: "Geist",
                    data: geistData,
                    style: 'italic'
                },
            ],
        }
    )
  }

          