import Head from 'next/head';
import { useRouter } from 'next/router';
import styles from '../styles/Home.module.css';

export default function Home() {

  const router = useRouter()

  const start = () => {
    router.push("/level/level1")
  }

  return (
    <div className={styles.container}>
      <Head>
        <title>Welcome</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <div>
      <div className='level-container'>
        <div className='level'>
          <h2 style={{textAlign : "center"}}>Hello! Enjoy the journey (hehehe)</h2>
          <button onClick={() => start()}>Start</button>
        </div>
      </div>
      </div>
    </div>
  )
}