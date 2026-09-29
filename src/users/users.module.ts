import { Module } from "@nestjs/common";
import { CreateUserUseCase } from "../adapters/postgres/usecase/create-user.usecase";

@Module({
	imports: [],
	providers: [CreateUserUseCase],
	controllers: []
})
export class UsersModule {
	
}
