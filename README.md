# Catalk

 A web app where you can chat with AI cats. Each cat has its own breed and personality. Ask questions, have conversations, or just say hi to your favorite virtual cat.

## Features

- **16 Cat Breeds** - Chat with Maine Coon, Siamese, Persian, Bengal, Sphynx, and more
- **AI Powered** - Uses Google's Gemini AI to make cats feel real
- **Google Login** - Sign in with your Google account
- **Save Chats** - Your chat history is saved

## Cat Breeds Available

- Maine Coon - Gentle and friendly giants
- Siamese - Vocal and social
- Persian - Calm and laid-back
- Bengal - Athletic and playful
- Sphynx - Affectionate and outgoing
- British Shorthair - Easygoing and calm
- Scottish Fold - Sweet and adaptable
- Abyssinian - Energetic and curious
- Burmese - Playful and social
- Bombay - Affectionate and vocal
- Japanese Bobtail - Playful and intelligent
- And more...

## How to Run

### You Need

- Go 1.22 or newer
- Node.js and npm
- PostgreSQL (we have Docker for this)
- Google API key for Gemini AI

### Setup

1. Copy the config file and add your settings:
```bash
cp config/config.example.yaml config/config.yaml
```

2. Add your Google API key to the config file

3. Start the database:
```bash
make docker-run
```

4. Run database migrations:
```bash
make migrate-up
```

5. Start the backend:
```bash
make run
```

6. In another window, start the web frontend:
```bash
cd cmd/web
npm install
npm run dev
```

7. Open your browser and go to `http://localhost:4321`

## Project Structure

```
catalk/
├── cmd/
│   ├── api/          # Go backend server
│   └── web/          # Astro frontend
├── internal/
│   ├── ai/           # Gemini AI integration
│   ├── auth/         # Google OAuth login
│   ├── database/     # Database connection
│   ├── server/       # HTTP routes and handlers
│   └── users/        # User management
├── instructions/     # AI instructions for each cat breed
└── config/          # Configuration files
```

## Built With

- **Backend**: Go
- **Frontend**: Astro + React + Tailwind CSS
- **Database**: PostgreSQL
- **AI**: Google Gemini
- **Auth**: Google OAuth2
