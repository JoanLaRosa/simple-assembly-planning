# Assembly Time Calculator

This program calculates the minimum time required to assemble a product with parallel assembly constraints.

## Problem Description

Given a tree-like assembly hierarchy where:

- Each node represents a component
- Components have assembly times and parallel assembly constraints
- Components may contain sub-components
- Parallel assembly is limited by max_parallel value

The program computes the minimum total assembly time considering parallel assembly opportunities.

## Example Structure

```json
{
  "name": "Product",
  "assembly_time": 5,
  "max_parallel": 2,
  "children": [
    {
      "name": "Component A",
      "assembly_time": 3,
      "max_parallel": 1,
      "children": [
        {
          "name": "Part A1",
          "assembly_time": 2,
          "max_parallel": 1,
          "children": []
        },
        {
          "name": "Part A2",
          "assembly_time": 4,
          "max_parallel": 1,
          "children": []
        }
      ]
    },
    {
      "name": "Component B",
      "assembly_time": 6,
      "max_parallel": 2,
      "children": [
        {
          "name": "Part B1",
          "assembly_time": 2,
          "max_parallel": 1,
          "children": []
        },
        {
          "name": "Part B2",
          "assembly_time": 1,
          "max_parallel": 1,
          "children": []
        }
      ]
    }
  ]
}
```

## Expected Output

Minimum Assembly Time: 11 minutes
