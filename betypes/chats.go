package betypes

// NewChats return *Chats.
//
// In goroutine starts adding users to chats from Queue.
func NewChats(usersCount, buffer int) *Chats {
	c := &Chats{
		Chats:      make([]Chat, 0),
		UsersCount: usersCount,
		Queue:      make(chan User, buffer),
	}

	go c.StartAddingUsers()

	return c
}

// StartAddingUsers adds users to chats form Queue.
//
// Recommended run in goroutine.
func (c *Chats) StartAddingUsers() {
	for u := range c.Queue {
		if len(c.Chats) == 0 {
			c.AddChat(*NewChat())
			c.Chats[len(c.Chats)-1].AddUser(u)
			continue
		}

		chats := c.FindSuitableChatsForUser(u)
		if chats != nil {
			chats[0].AddUser(u)
			continue
		}

		c.AddChat(*NewChat())
		c.Chats[len(c.Chats)-1].AddUser(u)
	}
}

// AddChat adds the Chat to []Chat.
func (c *Chats) AddChat(chat Chat) {
	c.Chats = append(c.Chats, chat)
}

// FindSuitableChatsForUser returns all suitable chats for the User.
//
// Takes into account the user's age, city, age of the interlocutors and the city of the user.
func (c *Chats) FindSuitableChatsForUser(user User) []*Chat {
	chats := make([]*Chat, 0)
	for i := 0; i < len(c.Chats); i++ {
		usersCount := len(c.Chats[i].Users)
		if usersCount == c.UsersCount {
			continue
		}

		u := c.Chats[i].Users[len(c.Chats[i].Users)-1]
		if user.IsSuitableAge(u) && user.IsSuitableCity(u) && user.IsSuitableSex(u) {
			chats = append(chats, &c.Chats[i])
		}
	}

	if len(chats) != 0 {
		return chats
	}

	return nil
}

// AddUserToQueue adds user to Queue.
func (c *Chats) AddUserToQueue(user User) {
	c.Queue <- user
}

// DeleteUserFromChat removes the user from the chat.
func (c *Chats) DeleteUserFromChat(userID int) {
	for i := 0; i < len(c.Chats); i++ {
		for j := 0; j < len(c.Chats[i].Users); j++ {
			if c.Chats[i].Users[j].ID == userID {
				c.Chats[i].Users[len(c.Chats[i].Users)-1],
					c.Chats[i].Users[i] = c.Chats[i].Users[i], c.Chats[i].Users[len(c.Chats[i].Users)-1]
				c.Chats[i].Users = c.Chats[i].Users[:len(c.Chats[i].Users)-1]
				return
			}
		}
	}
}

// DeleteChatWithUser delete the chat in which there is a User.
func (c *Chats) DeleteChatWithUser(userID int) {
	for i := 0; i < len(c.Chats); i++ {
		for j := 0; j < len(c.Chats[i].Users); j++ {
			if c.Chats[i].Users[j].ID == userID {
				c.Chats[len(c.Chats)-1], c.Chats[i] = c.Chats[i], c.Chats[len(c.Chats)-1]
				c.Chats = c.Chats[:len(c.Chats)-1]
				return
			}
		}
	}
}

// GetUserInterlocutors returns User interlocutors if User is in Chat.
func (c *Chats) GetUserInterlocutors(userID int) []User {
	if !c.IsUserInChat(userID) {
		return nil
	}

	for i := 0; i < len(c.Chats); i++ {
		if c.Chats[i].IsUserInChat(userID) {
			return c.Chats[i].GetUserInterlocutors(userID)
		}
	}

	return nil
}

// IsUserInChat returns true if the user is in chat.
func (c *Chats) IsUserInChat(userID int) bool {
	for i := 0; i < len(c.Chats); i++ {
		if c.Chats[i].IsUserInChat(userID) {
			return true
		}
	}
	return false
}
