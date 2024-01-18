import React from 'react'
import About from '../../../components/about/About'
import Mission from '../../../components/mission/Mission'
import Hero from '../../../components/hero/Hero'
import Spacer from '../../../components/spacer/Spacer'

const page = () => {
  return (
    <div>
      <Hero />
      <Spacer />
      <About />
      <Spacer />
      <Mission />
      <Spacer />
    </div>
  )
}

export default page