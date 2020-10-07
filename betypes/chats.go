package betypes

func (c *Chat) AddUserToChat(user User) {
	c.Users = append(c.Users, user)
}

func (c *Chats) AddUserToQueue(user User) {
	c.UsersQueue <- user
}

func (c *Chats) GetChatByID(chatID uint64) *Chat {
	for _, chat := range c.Chats {
		if chat.ID == chatID {
			return &chat
		}
	}
	return nil
}
