'use client'
import { createWeb3Modal, defaultConfig } from '@web3modal/ethers/react'

const projectId: any = process.env.PROJECT_ID

let api_key = process.env.API_KEY
let baseUrl = "https://base-mainnet.g.alchemy.com/v2/" + api_key;

const mainnet = {
    chainId: 8453,
    name: 'Base Mainnet',
    currency: 'ETH',
    explorerUrl: 'https://basescan.org',
    rpcUrl: baseUrl
  }

const metadata = {
    name: 'Versus',
    description: 'Create matches directly on farcaster with Versus, a bot powered by Stadium 🏟️',
    url: 'https://wag3r-bot.vercel.app/redirect', // origin must match your domain & subdomain
    icons: ['https://magenta-hollow-tiglon-795.mypinata.cloud/ipfs/QmWLS8z2ZcQ8JDGLZB9XB4LwXYbnmEQaJiG7fMyvuA42qP']
  }
  

  createWeb3Modal({
    ethersConfig: defaultConfig({ metadata }),
    chains: [mainnet],
    projectId,
    enableAnalytics: true // Optional - defaults to your Cloud configuration
  })
  
  export function Web3ModalProvider({ children }: any ) {
    return children
  }