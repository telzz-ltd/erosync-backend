import { Injectable } from '@nestjs/common';
import { Pool } from 'pg';
import { User } from '../../users/entities/user.entity';
import { UserRepository } from '../../users/ports/user-repository.port';

@Injectable()
export class PGUserRepository implements UserRepository {
  constructor(private readonly pool: Pool) {}
  async save(user: User): Promise<void> {
    await this.pool.query(
      `
			INSERT INTO users (id, name, email, password_hash, status, role, created_at, updated_at, deleted_at, email_verified_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
			ON CONFLICT (id) DO UPDATE SET
				name=EXCLUDED.name,
				email=EXCLUDED.email,
				password_hash=EXCLUDED.password_hash,
				status=EXCLUDED.status,
				role=EXCLUDED.role,
				updated_at=EXCLUDED.updated_at,
				deleted_at=EXCLUDED.deleted_at,
				email_verified_at=EXCLUDED.email_verified_at;
		`,
      [
        user.id,
        user.name,
        user.email,
        user.passwordHash,
        user.status,
        user.role,
        user.createdAt,
        user.updatedAt,
        user.deletedAt,
        user.emailVerifiedAt,
      ],
    );
  }

  async findById(id: string): Promise<User | null> {
    const result = await this.pool.query<User>(
      'SELECT * FROM users WHERE id = $1 LIMIT 1;',
      [id],
    );
    return result.rows[0] ?? null;
  }

  async findByEmail(email: string): Promise<User | null> {
    const result = await this.pool.query<User>(
      'SELECT * FROM users WHERE email = $1 LIMIT 1;',
      [email],
    );
    return result.rows[0] ?? null;
  }

  async existByEmail(email: string): Promise<boolean> {
    const result = await this.pool.query<{ count: number }>(
      'SELECT count(*) FROM users WHERE email = $1;',
      [email],
    );
    return result.rows[0].count > 0;
  }

  async delete(id: string): Promise<void> {
    await this.pool.query('DELETE FROM users WHERE id = $1;', [id]);
  }
}
