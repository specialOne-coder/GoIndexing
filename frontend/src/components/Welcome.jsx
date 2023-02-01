import React, { useContext } from 'react'
import { FaConnectdevelop } from 'react-icons/fa'
//import { AppContext } from '../context/AppContext';
import { useWeb3Modal } from '@web3modal/react'
import { shortenAddress } from '../utils/ShortAddress'

const Welcome = () => {
  //const { currentAccount, connectWallet } = useContext(AppContext);

  return (
    <div className="welcome flex max-w-[1500px] m-auto justify-center items-center p-[100px] ">
      <div className="welcome-div-text md:flex-[0.5] flex justify-center px-20 flex-wrap items-center self-center">
        <div className="flex flex-1 justify-center items-center flex-col  mf:ml-30">
          <h1 className="text-center sm:text-5xl text-white text-gradient py-1">
            Indexation
          </h1>
          <p className="text-center mt-5 text-white font-light md:w-9/12 w-11/12 text-base">
           Index your token history, download it and share it with your friends.
          </p>
        </div>
        <div className="welcome-button flex items-center cursor-pointer">
       
        </div>
      </div>
    </div>
  )
}

export default Welcome
