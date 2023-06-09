import sql from 'mssql';
import dbConfig from '../../constants/dbconfig';

export default async function login(req, res) {
  if (req.method === 'POST') {
    await sql.connect(dbConfig);

    const { identifier } = req.body;

    // Retrieve user from database
    //! chnage this to PUB_IDENTIFIER
    const user = await sql.query(`SELECT * FROM PLAYERS WHERE PLAYER_ID = '${identifier}'`);

    if (user.rowsAffected.length === 0) {
      return res.status(401).json({ message: 'Invalid credentials' });
    }

    return res.status(200).json({ token : user.recordset[0].PLAYER_ID });
  } else {
    res.status(405).json({ message: 'Method not allowed' });
  }
}
