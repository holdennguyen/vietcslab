Key Features:

Browse & Filter Courses
Purchase Courses using Stripe
Mark Chapters as Completed or Uncompleted
Progress Calculation of each Course
Student Dashboard
Teacher mode
Create new Courses
Create new Chapters
Easily reorder chapter position with drag n’ drop
Upload thumbnails, attachments and videos using UploadThing
Video processing using Mux
HLS Video player using Mux
Rich text editor for chapter description
Authentication using Clerk
ORM using Prisma
MySQL database using Planetscale

## Getting Started

First, run the development server:

```bash
npm run dev
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

## Prisma (MongoDB's ORM)

```bash
npx prisma studio
```

Open [http://localhost:5555](http://localhost:555) with your browser to see the collections.

## Test Stripe Integration in local

1. Download the Stripe CLI and log in with your Stripe account

    ```bash
    stripe login
    ```

2. Forward events to your destination

    ```bash
    stripe listen --forward-to localhost:3000/api/webhook
    ```

3. Trigger events with the CLI

    ```bash
    stripe trigger payment_intent.succeeded
    ```

## Deploy on Vercel

The easiest way to deploy your Next.js app is to use the [Vercel Platform](https://vercel.com/new?utm_medium=default-template&filter=next.js&utm_source=create-next-app&utm_campaign=create-next-app-readme) from the creators of Next.js.

Check out our [Next.js deployment documentation](https://nextjs.org/docs/deployment) for more details.
