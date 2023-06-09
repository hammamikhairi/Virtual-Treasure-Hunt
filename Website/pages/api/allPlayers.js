import sql from 'mssql';
import dbConfig from '../../constants/dbconfig';


export default async function handler(req, res) {
  try {
    // Connect to the database
    await sql.connect(dbConfig);

    // Query the database for all players from the 'PLAYERS' table
    const result = await sql.query('SELECT * FROM PLAYERS');

    // Send the players data as the response
    res.status(200).json(result.recordset);
  } catch (err) {
    console.error(err);
    res.status(500).send('Server error');
  } finally {
    // Close the database connection
    await sql.close();
  }
}


