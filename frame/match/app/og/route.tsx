/* eslint-disable @next/next/no-img-element  */
/* eslint-disable jsx-ally/alt-text */
// @ts-nocheck
import { ImageResponse } from 'next/server'

export const runtime = 'edge';


export async function GET(req: Request) {
  const imageData = await fetch(new URL('../../public/base.png', import.meta.url)).then(
    (res) => res.arrayBuffer(),
  );
  return new ImageResponse (
        (
          {staking()}
        )
    )
  }


  function staking() {
      return (
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
      
//   <div style={{position: 'relative', display: 'flex'}}>
//   <img 
//       src={imageData}
//    />
//   <div style={{display: 'flex', position: 'absolute', top: '65%', left: '50%', transform: 'translate(-50%, -50%)', textAlign: 'center', color: '#66757F', fontSize: '72px', fontWeight: '1200'}}> 
//           <span style={{fontStyle: 'italic'}}>STAKE $ 20 USDC</span>
//   </div>
// </div>