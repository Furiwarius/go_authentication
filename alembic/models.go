package alembic

type Account struct {
	id            int
	guid          [16]byte // guid переданный пользователем
	ip            string   // ip адрес пользователя
	email         string   // Почта пользователя
	refresh_token string   // Токен для обновления Access токена
}
