# Use the official Node.js image as the base image
FROM node:16-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the entire project directory into the container
COPY . .

# Build the frontend app
RUN npm run build

# Install serve to serve the production build
RUN npm install -g serve

# Expose port 3000 to the outside world
EXPOSE 3000

# Command to serve the built app
CMD ["serve", "-s", "build"]
