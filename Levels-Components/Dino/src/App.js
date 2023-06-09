import Game from './lib';

const App = () => (
  <>
  <Game />
    <div style={{ display : "flex", alignItems: "center", flexDirection : "column"}}>
      <div style={{width : "600px"}}>

      <h1>No internet</h1>
      Try:
      <ul>
        <li>

        Checking the network cables, modem, and router
        </li>
        <li>
        Reconnecting to Wi-Fi

        </li>
      </ul>
        ERR_INTERNET_DISCONNECTED
      </div>
    </div>
  </>
);

export default App;
