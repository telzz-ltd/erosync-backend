export class ResponseMapper {
	static success(data: any){
		return {"message": "success", data};
	}
}
