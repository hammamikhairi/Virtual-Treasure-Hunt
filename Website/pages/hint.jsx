import dynamic from 'next/dynamic';

const ChromeDinoComponentNoSsr = dynamic(() => import('../componenets/dino/Dino'), {
  ssr: false,
});

function HomePage() {
  return (
    <div>
      <ChromeDinoComponentNoSsr />
      <div style={{ display : "flex", alignItems: "center", flexDirection : "column"}}>
      <div style={{width : "600px"}}>

      <h1>Internal Server Error</h1>
      Try:
      <ul>
        <li>
          Refreshing The Page
        </li>
        <li>
         That won't work, Call 56-488-593
        </li>
      </ul>
        SERVER_ERR_OR_IS_IT_?
      </div>
    </div>

    </div>
  );
}

export default HomePage;
