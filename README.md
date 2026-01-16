# Restaurant Management System

A full-stack restaurant management application designed to manage menu items and customer orders efficiently.  
This repository is intended primarily for **developers** who want to understand, extend, or build upon a practical full-stack project using **React**, **Node.js**, and **MySQL**.


## Features

- **Menu Management**
  - Add, update, and delete menu items
  - Categorize items with descriptions and prices

- **Order Management**
  - Create customer orders
  - Track order status (pending, preparing, delivered)
  - View order details with associated items

- **Real-Time Updates**
  - Frontend reflects backend changes without manual refresh

- **Customer Handling**
  - Basic customer identification per order


## Tech Stack

### Frontend
- React
- Vite
- Tailwind CSS
- JavaScript

### Backend
- Node.js
- Express.js
- MySQL

### Database
- MySQL 8  
- Tested with Azure Database for MySQL (Flexible Server)


## Project Structure

```plaintext
restaurant-management/
├── frontend/
│   ├── src/
│   ├── public/
│   ├── package.json
│   └── vite.config.js
├── backend/
│   ├── server.js
│   ├── package.json
│   └── .env.example
├── README.md
└── LICENSE
```


## Installation and Setup

### Prerequisites

- Node.js (v14 or higher)
- npm or yarn
- MySQL
- Git


## Backend Setup

```bash
git clone https://github.com/duronto23/restaurant-management.git
cd restaurant-management/backend
npm install
```

Create a `.env` file:

```env
AZURE_MYSQL_HOST=localhost
AZURE_MYSQL_USER=root
AZURE_MYSQL_PASSWORD=your_password
AZURE_MYSQL_DATABASE=restaurant_db
```

Run backend:
```bash
npm run dev
```


## Frontend Setup

```bash
cd ../frontend
npm install
npm run dev
```


## API Endpoints

| Method | Endpoint | Description |
|------|---------|-------------|
| GET | /api/menu-items | Fetch all menu items |
| POST | /api/menu-items | Create a new menu item |
| GET | /api/orders | Fetch all orders |
| POST | /api/orders | Create a new order |
| PUT | /api/orders/:id | Update order status |

