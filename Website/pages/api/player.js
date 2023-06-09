import sql from 'mssql';
import dbConfig from '../../constants/dbconfig';


export default async function handler(req, res) {
  try {
    // Connect to the database
    await sql.connect(dbConfig);

    // Get the player ID from the query parameters
    const playerId = req.query.id;

    // Query the database for the player with the given ID
    const result = await sql.query`SELECT * FROM PLAYERS WHERE PLAYER_ID = ${playerId}`;

    // If the player is found, send the player data as the response
    if (result.recordset.length > 0) {
      res.status(200).json(result.recordset[0]);
    } else {
      res.status(404).send('Player not found');
    }
  } catch (err) {
    console.error(err);
    res.status(500).send('Server error');
  } finally {
    // Close the database connection
    await sql.close();
  }
}
