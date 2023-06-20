import Head from 'next/head';
import { useRouter } from 'next/router';
import { useEffect, useState } from 'react';
import DownloadButton from '../../componenets/DownloadButton/DownloadButton';
import FlagForm from '../../componenets/FlagForm/FlagForm';

const BASE_URL = "http://localhost:5051/"
const BASE_WEBSITE_URL = "http://localhost:3000/"


export async function getServerSideProps(context) {

  const {id} = context.params
  const currentLevel = id.match(/\d+$/)[0]


  if (currentLevel == 10) {
    return  {
      redirect: {
        destination : "/congrats",
        permanent: false,
      }
    }
  }

  const levelRes = await fetch(`${BASE_URL}getLevel?level=${currentLevel}`)
  const level = await levelRes.json()

  return {
    props: {
      level,
      currentLevel
    }
  };

}

function Level({ level, currentLevel }) {
  const router = useRouter();
  const [hidden, setHidden] = useState(level.Message.includes("$$") ? level.Message.split("$$")[1] : '')
  const levelMessage = level.Message.split("$$")[0];

  useEffect(() => {
    if (!level.Message.includes("$$")) {
      setHidden('')
    }
  }, [router.asPath])

  return (
    <>
      <Head>
        <title>{level.Title}</title>
        {
          hidden != '' &&
          <noscript>{hidden}</noscript>
        }
      </Head>
      <div className='level-container'>
        <div className='level'>
          <h2 style={{textAlign : "center"}}>{levelMessage}</h2>
          {
            currentLevel == 4 &&
            <a href={`${BASE_WEBSITE_URL}hint`} target="_blank">hint</a>
          }
          {
            level.Images && !level.Images[0].includes(".mp3") &&
            <img className='level-image' src={`/images/${level.Images[0]}`} alt="idk" />
          }
          {
            level.Audio &&
            <audio controls>
              <source src={`/assets/${level.Audio}`} type="audio/mpeg" />
              Browser doesn't support audio files
            </audio>
          }
          {
            level.Files &&
            <DownloadButton level={currentLevel} file={level.Files[0]} />
          }
          <FlagForm level={currentLevel} />
        </div>
      </div>
    </>
  );
}

export default Level;
