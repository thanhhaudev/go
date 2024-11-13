```text
+--------------+          +--------------+          +-------------+
|  Dispatcher  |          |   Worker 1   |          |   Worker 2  |          
+--------------+          +--------------+          +-------------+
|                           |                           |
| Dispatch Job 1            |                           |
+-------------------------> |                           |
|                           | Process Job 1             |
|                           +-------------------------> |
| Dispatch Job 2            |                           |
+-----------------------------------------------------> |
|                           |                           | Process Job 2
| Wait for Completion       | Finish Job 1              |
|                           |                           |
|                           | Register in Pool          |
| Dispatch Job 3            |                           |
+-------------------------> |                           |
|                           | Process Job 3             |
|                           +-------------------------> |
|                           |                           | Finish Job 2
|                           |                           |
| Wait for Completion       |                           |
|                           |                           | Register in Pool
| Dispatch Job 4            |                           |
+-----------------------------------------------------> |
|                           |                           | Process Job 4
```
### Activity Diagram Description:

- **Job 1** is assigned to **Worker 1**, and **Job 2** is assigned to **Worker 2**.
- **Worker 1** completes **Job 1** and returns to the **Pool**, then takes on **Job 3**.
- **Worker 2** completes **Job 2** and returns to the **Pool** to be ready for **Job 4**.
- **Job 4** is assigned to **Worker 2**.

### Detailed Explanation:

- **Dynamic Assignment**: The Dispatcher will choose whichever **Worker** is available to assign the next job.
- **Load Balancing**: Ensures that no **Worker** is overloaded, as it always picks an idle **Worker**.
- **Reusability**: After completing a job, the **Worker** re-registers in the **work pool** and waits for new jobs.

This diagram illustrates the dynamic process of job assignment to **Workers** based on the **work pool**.