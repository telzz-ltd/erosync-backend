import { Body, Controller, Post } from "@nestjs/common";
import { CreateUserUseCase } from "../adapters/postgres/usecase/create-user.usecase";
import { ResponseMapper } from "../utils/response-mapper";
import { RegisterRequest } from "./dto/register";

@Controller()
export class UsersController {
	constructor(
		private createUser: CreateUserUseCase
	){}

	@Post("/auth/register")
	async registerUser(@Body() req: RegisterRequest){
		const result = this.createUser.execute(req);
		return ResponseMapper.success(result);
	}
}
