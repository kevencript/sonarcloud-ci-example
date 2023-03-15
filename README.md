# Continuous Integration (CI) with GitActions & [Sonarcloud](https://sonarcloud.io) - Pushing covered Go apps to Dockerhub 

![](https://assets-eu-01.kc-usercontent.com/45f00125-6dea-0121-0efe-ce8937882537/7fd02d4d-94bc-4d88-a77e-db2d6e81a658/body-3cedacd4-0d81-4da8-986c-7978d6765e79_SC-SQ%2BBlog%2B-%2Bbranding%25402x.png?w=1432&h=448&auto=format&fit=crop)

## Overview
This project is a pratical example of how to create a Continuous Integration (CI) to perform code coverage, analysis and quality checks automatically with [Sonarcloud](https://sonarcloud.io) & Github Actions. This integration enables you to automatically run code analysis using SonarCloud and to provide feedback on the quality of your code directly in your GitHub repository.

For this project we are deploying an illustrative example Go app through CI/CD workflows to Dockerhub.

## Continuous Integration (CI) Workflow (GitActions)

* :open_file_folder: Checkout the code from your repository
* :hammer_and_wrench: Set up your build environment (such as installing dependencies)
* :microscope: Run your build and test scripts
* :mag: Run the SonarCloud scanner to analyze your code
* :chart_with_upwards_trend: Publish the analysis results to SonarCloud

## What is "Sonarqube"? 

[SonarQube](https://sonarcloud.io) is an open-source platform for continuous code quality inspection and analysis. It provides detailed reports on code issues and allows developers to track and improve the quality of their code over time. With SonarQube, developers can quickly identify and fix issues in their code, improving the maintainability and reliability of their applications. Additionally, it can be integrated with popular development tools for automatic code analysis and instant feedback to developers.

#### Example of a code coverage interface on Sonarqube
![](https://assets-eu-01.kc-usercontent.com/e1f3885c-805a-0150-804f-0996e00cd37d/aeb9bf5a-a2bd-4b70-98b3-3499e7c796b3/successfulproject.png?w=813&h=502&auto=format&fit=crop)
> Here we can see how powerfull Sonarqube can be showing us the bugs, code smells and other important infos related the code


## Sonarqube usual platform therms
Here we have some concerns related to how we can create rules in which will define what is a bug or something critical, quality profiles and the quality gates, in which will define the quality standarts for our analysis.

### Sonarqube: Quality Profiles 👥
Quality Profiles are a key component of SonarCloud that allows you to define a set of rules to enforce code quality standards across your projects. You can choose from a variety of pre-defined profiles or create your own custom profile. The profile can be applied to individual projects or across your organization.

### Sonarqube: Rules 📜
SonarCloud provides a set of pre-defined rules for different programming languages. These rules check for various code quality issues, such as complexity, duplications, and security vulnerabilities. You can also create custom rules to meet specific requirements. The rules are applied to the code during analysis and reported in the SonarCloud dashboard.

### Quality Gates 🚦
Quality Gates are used to define the quality standards that must be met before code can be deployed to production. The Quality Gates in SonarCloud can be customized to include specific quality checks, such as code coverage, maintainability, and security. If the code does not meet the quality standards defined in the Quality Gate, the build will fail, and the issues will be reported in the SonarCloud dashboard.

Overall, these three components work together to ensure that your code meets the highest quality standards and helps you identify potential issues before they become a problem. With SonarCloud, you can create custom Quality Profiles, define rules to enforce standards, and set Quality Gates to ensure that your code is ready for deployment.
