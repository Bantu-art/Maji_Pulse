# Water Management System Dashboard

This project is a Water Management System Dashboard designed for both admin and user roles. It simulates real-time data for monitoring water flow rates, detecting leakages, and ensuring fair distribution of water in different areas. The project is built using Go for the backend and HTML, CSS, and JavaScript for the frontend.

## Features
- **Admin Dashboard**:
  - Displays flow rates for different areas.
  - Detects and alerts for water leakages.
  - Shows fair distribution percentages for each area.

- **User Dashboard**:
  - Displays real-time flow rates.
  - Shows the cost associated with the water usage.

- **Leakage Simulation**:
  - Simulates leakages that start after 1 minute and stop after 2 minutes to demonstrate leak detection and repair.

## Project Structure
```bash

├── main.go
├── templates
│ ├── admin.html
│ ├── user.html
│ ├── index.html
│ └── signup.html
├── static
│ ├── admin.js
│ ├── user.js
│ └── styles.css
├── README.md
└── go.mod

```


## Getting Started

### Prerequisites

- Go 1.16 or later
- A modern web browser

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/water-management-system.git
   cd water-management-system
   ```

2. Install dependencies:
```bash
go mod tidy

```
## Running the Application

1.  Start the Go server:
```bash
go run main.go

```
2. Open your web browser and navigate to:
```bash

http://localhost:8060/ for the homepage
http://localhost:8060/admin for the admin dashboard
http://localhost:8060/user for the user dashboard
http://localhost:8060/signup for the signup page
```

## Project Details

main.go: Contains the main application logic, including HTTP handlers and data simulation.
templates: Contains the HTML templates for different pages.
static: Contains the static files such as JavaScript and CSS.



## Data Simulation

 Admin Data: Simulates flow rates, leakages, and fair distribution percentages for different areas.
 User Data: Simulates flow rates and costs for users.


## Leakage Simulation


Leakage starts after 1 minute and stops after 2 minutes, simulating a real-time detection and repair system. The flow rate fluctuates to reflect leakage conditions, and the leakage status is updated accordingly.


## Contributing

Contributions are welcome! Please fork the repository and submit pull requests.

## License

This project is licensed under the MIT License.
