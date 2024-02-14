import { defaultWagmiConfig } from '@web3modal/wagmi/react/config'
import { cookieStorage, createStorage } from 'wagmi'
import { polygon } from 'wagmi/chains'

export const projectId = process.env.NEXT_PUBLIC_PROJECT_ID

if (!projectId) {
  throw new Error('Project ID is not defined')
}

const metadata = {
  name: 'Versus',
  description: 'Create matches directly on farcaster with Versus, a bot powered by Stadium 🏟️',
  url: 'https://wag3r-bot.vercel.app/', // origin must match your domain & subdomain
  icons: ['https://magenta-hollow-tiglon-795.mypinata.cloud/ipfs/QmWLS8z2ZcQ8JDGLZB9XB4LwXYbnmEQaJiG7fMyvuA42qP']
}

export const config = defaultWagmiConfig({
  projectId,
  chains: [polygon],
  metadata: metadata,
  enableInjected: true,
  enableWalletConnect: true,
  enableEIP6963: true,
  enableCoinbase: true,
  enableEmail: true,
  storage: createStorage({
    storage: cookieStorage
  }),
  ssr: true
})






// let api_key = process.env.API_KEY
// if (!api_key) {
//   throw new Error('Project ID is not defined')
// }
// let baseUrl = "https://base-mainnet.g.alchemy.com/v2/" + api_key;

// const mainnet = {
//     chainId: 8453,
//     name: 'Base Mainnet',
//     currency: 'ETH',
//     explorerUrl: 'https://basescan.org',
//     rpcUrl: baseUrl
//   }



//   createWeb3Modal({
//     ethersConfig: defaultConfig({ metadata }),
//     chains: [mainnet],
//     projectId,
//     enableAnalytics: true // Optional - defaults to your Cloud configuration
//   })
  
//   export function Web3ModalProvider({ children }: any ) {
//     return children
//   }