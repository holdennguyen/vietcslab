# Key Features

- Browse & Filter Courses
- Purchase Courses using Stripe
- Mark Chapters as Completed or Uncompleted
- Progress Calculation of each Course
- Student Dashboard
- Teacher mode
- Create new Courses
- Create new Chapters
- Easily reorder chapter position with drag n’ drop
- Upload thumbnails, attachments and videos using UploadThing
- Video processing using Mux
- HLS Video player using Mux
- Rich text editor for chapter description
- Authentication using Clerk
- ORM using Prisma

## Getting Started

First, run the development server:

```bash
npm run dev
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

## Prisma (MongoDB's ORM)

To generate the Prisma client and push the schema to the database, run:

```bash
npx prisma generate
npx prisma db push
```

To run the seed script, use:

```bash
node scripts/seed.ts
```

To open Prisma Studio, run:

```bash
npx prisma studio
```

Open [http://localhost:5555](http://localhost:5555) with your browser to see the collections.

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

1. Add the following to the `scripts` section of your `package.json` file:

    ```json
    "postinstall": "prisma generate"
    ```

2. Integrate your GitHub repository with your Vercel project and update the `.env` file to the Environment Variables of the Vercel project. Deploy the project.

3. After the first successful deployment, update the `NEXT_PUBLIC_APP_URL` environment variable with your Vercel app domain.

4. Configure Stripe Test mode:
    - Go to the Stripe developer dashboard, navigate to the Event destinations tab.
    - Add a destination from your account and select all Checkout events.
    - Input your Vercel app URL as the Webhook endpoint (e.g., `https://*.vercel.app/api/webhook`).
    - Copy the Signing secret and update the `STRIPE_WEBHOOK_SECRET` environment variable in your Vercel app.

5. Re-deploy the Vercel app without using the existing build cache.

> Use the test credit card number `4242 4242 4242 4242` to simulate a successful payment in Stripe.
