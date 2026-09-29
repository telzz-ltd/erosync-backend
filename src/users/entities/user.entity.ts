export enum UserRole {
	USER = "USER",
	ADMIN = "ADMIN",
	STAFF = "STAFF"
}

export enum UserStatus {
	ACTIVE = "ACTIVE",
	INACTIVE = "INACTIVE"
}

export class User {
	id!: string;
	name!: string;
	email!: string;
	passwordHash!: string;
	role!: UserRole;
	status!: UserStatus;
	createdAt!: Date;
	updatedAt!: Date;
	deletedAt?: Date|null;
	emailVerifiedAt!: Date;
}
