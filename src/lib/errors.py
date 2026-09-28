
class FieldExistException(Exception):
    def __init__(self, field: str) -> None:
        super().__init__(f"{field} already exists")


class RecordNotFoundException(Exception):
    def __init__(self) -> None:
        super().__init__("record not found")
