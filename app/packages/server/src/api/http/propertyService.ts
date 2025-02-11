import pool  from "./database";


const getAllProperties = async () => {
    const query = 'SELECT * FROM properties';
    // const values = [userId];
    try {
      const result = await pool.query(query);
      return result || null;
    } catch (error) {
      console.error('Database query error:', error);
      throw error;
    }
  };

export { getAllProperties }
