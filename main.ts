/**
 * @author: Zachary Fowler
 * @version: 1.0.0
 * @date: 2025-11-20
 * @fileoverview: This file alculates the price of nuts, bolts, and washers
 */

// Assign constants
const boltPrice = 10;
const nutPrice = 5;
const washerPrice = 3;

// Input
let bolts = Number(prompt("Enter the number of bolts:"));
let nuts = Number(prompt("Enter the number of nuts:"));
let washers = Number(prompt("Enter the number of washers:"));

// Calculate total cost
let totalCost = (bolts * boltPrice) + (nuts * nutPrice) + (washers * washerPrice);

// Output order
console.log(`Number of bolts: ${bolts}`);
console.log(`Number of nuts: ${nuts}`);
console.log(`Number of washers: ${washers}`);

// Check order
if (nuts < bolts) {
  console.log("Check the Order, not enough nuts for the bolts you purchased.");
}
if (washers < bolts) {
  console.log("Check the Order, not enough washers for the bolts you purchased.");
}
if (nuts >= bolts && washers >= bolts) {
  console.log("Order is OK.");
}

// Output total cost
console.log("Your total cost of the order is " + totalCost + " cents.");

console.log("\nDone.")