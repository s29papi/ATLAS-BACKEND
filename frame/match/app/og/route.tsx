/* eslint-disable @next/next/no-img-element  */
/* eslint-disable jsx-ally/alt-text */
// @ts-nocheck
import { ImageResponse } from 'next/server'

export const runtime = 'edge';

const imageData = await fetch(new URL('../../public/base.png', import.meta.url)).then(
  (res) => res.arrayBuffer(),
);

export async function GET(req: Request) {

  return stakingTing()
  }

function stakingTing() {
  return new ImageResponse (
    <div style={{position: 'relative', display: 'flex'}}>
        <img 
            src={imageData}
         />
        <div style={{display: 'flex', position: 'absolute', top: '65%', left: '50%', transform: 'translate(-50%, -50%)', textAlign: 'center', color: '#66757F', fontSize: '72px', fontWeight: '1200'}}> 
                <span style={{fontStyle: 'italic'}}>STAKE $ 20 USDC</span>
        </div>
      </div>
)
}



// export async function GET(req: Request) {

//   return new ImageResponse (
//             <div style={{position: 'relative', display: 'flex'}}>
//                 <img 
//                     src={imageData}
//                  />
//                 <div style={{display: 'flex', position: 'absolute', top: '65%', left: '50%', transform: 'translate(-50%, -50%)', textAlign: 'center', color: '#66757F', fontSize: '72px', fontWeight: '1200'}}> 
//                         <span style={{fontStyle: 'italic'}}>STAKE $ 20 USDC</span>
//                 </div>
//               </div>
//     )
//   }
