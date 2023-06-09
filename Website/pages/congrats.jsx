import cookies from 'next-cookies';
import Head from 'next/head';
import { useRouter } from 'next/router';

const BASE_URL = "http://20.111.33.21/"
// const BASE_URL = "http://localhost:5051/"
const BASE_WEBSITE_URL = "http://localhost:3000/"


export async function getServerSideProps(context) {
  const { token } = cookies(context);

  if (!token && context.req.url !== '/login') {
    return {
      redirect: {
        destination: '/login',
        permanent: false,
      },
    };
  }
  const res = await fetch(`${BASE_URL}/player?playerId=${token}`);
  const user = await res.json();

  if (10 != user.level) {
    return {
      redirect: {
        destination: `/level/level${user.level}`,
        permanent: false,
      },
    };
  }

  return {
    props: {
      user
    }
  };
}



function Level({ user }) {
  const router = useRouter();
  
  return (
    <>
      <Head>
        <title>Congratulations!</title>
      </Head>
      <div className='level-container'>
        <div className='level'>
          <h2 style={{textAlign : "center"}}>congrats</h2>
        </div>
      </div>
    </>
  );
}

export default Level;
