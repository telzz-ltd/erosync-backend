import { Injectable } from '@nestjs/common';


type CreateUserCommand = {
	name: string;
	email: string;
	password: string;
}

@Injectable()
export class CreateUserUseCase {
	constructor() { }
	async execute(cmd: CreateUserCommand) {

	}
}
