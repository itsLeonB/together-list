const { Pool } = require('pg');
const { DATABASE } = require('../config');

class PostgresStore {
  constructor(sessionId = 'default') {
    this.sessionId = sessionId;
    this.pool = new Pool({
      connectionString: DATABASE.URL,
    });
    this.table = 'whatsapp_sessions';
  }

  async init() {
    await this.pool.query(`
      CREATE TABLE IF NOT EXISTS ${this.table} (
        session_id TEXT PRIMARY KEY
      );
    `);
  }

  // Check if a session exists
  async sessionExists({ session }) {
    const res = await this.pool.query(
      `SELECT 1 FROM ${this.table} WHERE session_id = $1 LIMIT 1`,
      [session]
    );
    return res.rowCount > 0;
  }

  // Delete session
  async delete({ session }) {
    await this.pool.query(`DELETE FROM ${this.table} WHERE session_id = $1`, [
      session,
    ]);
  }

  // Save full session object
  async save({ session }) {
    await this.pool.query(
      `INSERT INTO ${this.table} (session_id)
       VALUES ($1)
       ON CONFLICT (session_id)
       DO NOTHING`,
      [session]
    );
  }

  // Extract full session object
  async extract({ session, path }) {
    const res = await this.pool.query(
      `SELECT data FROM ${this.table} WHERE session_id = $1`,
      [session]
    );
    return res.rows[0]?.data || null;
  }
}

module.exports = PostgresStore;
