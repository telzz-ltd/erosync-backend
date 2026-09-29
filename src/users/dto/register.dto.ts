import { IsEmail, IsNotEmpty, Length, Matches } from 'class-validator';

export class RegisterRequest {
  @IsNotEmpty()
  @Matches('^[a-zA-Z]{2,}(?:( [a-zA-Z]{2,})){1,2}')
  name!: string;

  @IsNotEmpty()
  @IsEmail()
  email!: string;

  @IsNotEmpty()
  @Length(8, 50)
  password!: string;
}
