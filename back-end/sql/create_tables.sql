-- Create database if it doesn't exist
CREATE DATABASE IF NOT EXISTS food4student;
USE food4student;

-- Create tables (will be handled by GORM, but defining them here for completeness)
-- Tables will be created with appropriate plural names by GORM convention

-- Sample data for users
INSERT INTO users (id, created_at, updated_at, first_name, last_name, email, password, address, phone_number)
VALUES
    (1, NOW(), NOW(), 'John', 'Doe', 'john@example.com', '$2a$10$jlFYUIV4RZ4cfAiKt/QjrOVDRz491aITLthZkSEAJ/clIgve6DHpC', '123 College St', '555-1234'),  -- password: password123
    (2, NOW(), NOW(), 'Jane', 'Smith', 'jane@example.com', '$2a$10$jlFYUIV4RZ4cfAiKt/QjrOVDRz491aITLthZkSEAJ/clIgve6DHpC', '456 University Ave', '555-5678'),
    (3, NOW(), NOW(), 'Bob', 'Johnson', 'bob@example.com', '$2a$10$jlFYUIV4RZ4cfAiKt/QjrOVDRz491aITLthZkSEAJ/clIgve6DHpC', '789 Dorm Lane', '555-9012'),
    (4, NOW(), NOW(), 'Alice', 'Brown', 'alice@example.com', '$2a$10$jlFYUIV4RZ4cfAiKt/QjrOVDRz491aITLthZkSEAJ/clIgve6DHpC', '101 Campus Road', '555-3456');

-- Sample data for restaurants
INSERT INTO restaurants (id, created_at, updated_at, name, description, address, phone_number, status, logo_url, banner_url, total_rating, average_rating)
VALUES
    (1, NOW(), NOW(), 'Campus Pizza', 'Best pizza near campus', '100 College St', '555-1111', 4, 'https://example.com/logos/campus-pizza.jpg', 'https://example.com/banners/campus-pizza.jpg', 45, 4.5),
    (2, NOW(), NOW(), 'Burger Joint', 'Gourmet burgers for students', '200 University Blvd', '555-2222', 4, 'https://example.com/logos/burger-joint.jpg', 'https://example.com/banners/burger-joint.jpg', 35, 4.3),
    (3, NOW(), NOW(), 'Sushi Express', 'Fast Japanese cuisine', '300 Dorm Road', '555-3333', 4, 'https://example.com/logos/sushi-express.jpg', 'https://example.com/banners/sushi-express.jpg', 28, 4.0),
    (4, NOW(), NOW(), 'Taco Town', 'Authentic Mexican food', '400 Student Lane', '555-4444', 0, 'https://example.com/logos/taco-town.jpg', 'https://example.com/banners/taco-town.jpg', 0, 0.0),
    (5, NOW(), NOW(), 'Noodle House', 'Asian noodle specialties', '500 Academic Dr', '555-5555', 1, 'https://example.com/logos/noodle-house.jpg', 'https://example.com/banners/noodle-house.jpg', 0, 0.0);

-- Sample data for food_categories
INSERT INTO food_categories (id, created_at, updated_at, name, restaurant_id)
VALUES
    (1, NOW(), NOW(), 'Pizzas', 1),
    (2, NOW(), NOW(), 'Sides', 1),
    (3, NOW(), NOW(), 'Drinks', 1),
    (4, NOW(), NOW(), 'Burgers', 2),
    (5, NOW(), NOW(), 'Fries', 2),
    (6, NOW(), NOW(), 'Sushi Rolls', 3),
    (7, NOW(), NOW(), 'Bento Boxes', 3);

-- Sample data for food_items
INSERT INTO food_items (id, created_at, updated_at, name, description, base_price, photo_url, food_category_id)
VALUES
    (1, NOW(), NOW(), 'Pepperoni Pizza', 'Classic pepperoni pizza', '12.99', 'https://example.com/food/pepperoni.jpg', 1),
    (2, NOW(), NOW(), 'Cheese Pizza', 'Simple and delicious cheese pizza', '10.99', 'https://example.com/food/cheese.jpg', 1),
    (3, NOW(), NOW(), 'Garlic Breadsticks', 'Tasty garlic breadsticks', '4.99', 'https://example.com/food/breadsticks.jpg', 2),
    (4, NOW(), NOW(), 'Soda', 'Refreshing soft drinks', '1.99', 'https://example.com/food/soda.jpg', 3),
    (5, NOW(), NOW(), 'Classic Burger', 'Beef patty with lettuce, tomato, and cheese', '8.99', 'https://example.com/food/classic-burger.jpg', 4),
    (6, NOW(), NOW(), 'California Roll', 'Crab, avocado, cucumber', '7.99', 'https://example.com/food/california-roll.jpg', 6),
    (7, NOW(), NOW(), 'Chicken Teriyaki Bento', 'Teriyaki chicken with rice and vegetables', '11.99', 'https://example.com/food/teriyaki-bento.jpg', 7);

-- Sample data for variations
INSERT INTO variations (id, created_at, updated_at, name, min_select, max_select, food_item_id)
VALUES
    (1, NOW(), NOW(), 'Size', 1, 1, 1),
    (2, NOW(), NOW(), 'Extra Toppings', 0, 5, 1),
    (3, NOW(), NOW(), 'Size', 1, 1, 2),
    (4, NOW(), NOW(), 'Soda Type', 1, 1, 4),
    (5, NOW(), NOW(), 'Patty Options', 1, 1, 5),
    (6, NOW(), NOW(), 'Spice Level', 1, 1, 6);

-- Sample data for variation_options
INSERT INTO variation_options (id, created_at, updated_at, name, price_adjustment, variation_id)
VALUES
    (1, NOW(), NOW(), 'Small', -200, 1),
    (2, NOW(), NOW(), 'Medium', 0, 1),
    (3, NOW(), NOW(), 'Large', 300, 1),
    (4, NOW(), NOW(), 'Extra Cheese', 150, 2),
    (5, NOW(), NOW(), 'Mushrooms', 100, 2),
    (6, NOW(), NOW(), 'Peppers', 100, 2),
    (7, NOW(), NOW(), 'Small', -200, 3),
    (8, NOW(), NOW(), 'Medium', 0, 3),
    (9, NOW(), NOW(), 'Large', 300, 3),
    (10, NOW(), NOW(), 'Cola', 0, 4),
    (11, NOW(), NOW(), 'Lemon-Lime', 0, 4),
    (12, NOW(), NOW(), 'Root Beer', 0, 4),
    (13, NOW(), NOW(), 'Regular Beef', 0, 5),
    (14, NOW(), NOW(), 'Beyond Meat', 200, 5),
    (15, NOW(), NOW(), 'No Spice', 0, 6),
    (16, NOW(), NOW(), 'Medium Spice', 0, 6),
    (17, NOW(), NOW(), 'Extra Spicy', 50, 6);

-- Sample data for orders
INSERT INTO orders (id, created_at, updated_at, order_status, note, total_price, shipping_address, phone_number, orderer, user_id, restaurant_id)
VALUES
    (1, NOW(), NOW(), 3, 'Please deliver to front desk', 2899, '123 College St', '555-1234', 'John Doe', 1, 1),
    (2, NOW(), NOW(), 2, 'Ring bell when arrived', 1599, '456 University Ave', '555-5678', 'Jane Smith', 2, 2),
    (3, NOW(), NOW(), 1, 'Extra napkins please', 2399, '789 Dorm Lane', '555-9012', 'Bob Johnson', 3, 3),
    (4, NOW(), NOW(), 0, '', 1899, '101 Campus Road', '555-3456', 'Alice Brown', 4, 1);

-- Sample data for order_items
INSERT INTO order_items (id, created_at, updated_at, price, quantity, food_item_photo_url, variations, order_id)
VALUES
    (1, NOW(), NOW(), 1299, 2, 'https://example.com/food/pepperoni.jpg', 'Medium, Extra Cheese', 1),
    (2, NOW(), NOW(), 499, 1, 'https://example.com/food/breadsticks.jpg', '', 1),
    (3, NOW(), NOW(), 1599, 1, 'https://example.com/food/classic-burger.jpg', 'Beyond Meat patty', 2),
    (4, NOW(), NOW(), 799, 3, 'https://example.com/food/california-roll.jpg', 'Medium Spice', 3),
    (5, NOW(), NOW(), 1199, 1, 'https://example.com/food/pepperoni.jpg', 'Large', 4),
    (6, NOW(), NOW(), 199, 2, 'https://example.com/food/soda.jpg', 'Cola', 4);

-- Sample data for ratings
INSERT INTO ratings (id, created_at, updated_at, comment, stars, user_id, order_id)
VALUES
    (1, NOW(), NOW(), 'Pizza arrived hot and delicious!', 5, 1, 1),
    (2, NOW(), NOW(), 'Burger was good but delivery was slow', 3, 2, 2),
    (3, NOW(), NOW(), 'Best sushi near campus', 4, 3, 3);

-- Sample data for favorites
INSERT INTO favorites (id, created_at, updated_at, user_id, restaurant_id)
VALUES
    (1, NOW(), NOW(), 1, 1),
    (2, NOW(), NOW(), 1, 3),
    (3, NOW(), NOW(), 2, 2),
    (4, NOW(), NOW(), 3, 1),
    (5, NOW(), NOW(), 4, 3);