const { PrismaClient } = require("@prisma/client");

const database = new PrismaClient();

async function main() {
    try {
        await database.category.createMany({
            data: [
                { name: "Programming Fundamentals" },
                { name: "Software Development" },
                { name: "Web Development" },
                { name: "Data Science & Machine Learning" },
                { name: "Cybersecurity" },
                { name: "Operating Systems & Networking" },
                { name: "Computer Science Theory" },
                { name: "Career Guidance & Resources" },
            ]
        });
  
        console.log("Success");
    } catch (error) {
        console.log("Error sedding the database categories", error);
    } finally {
        await database.$disconnect();
    }
}

main();