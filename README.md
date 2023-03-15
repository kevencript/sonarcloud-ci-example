# Continuous Integration (CI) with GitActions & [Sonarcloud](https://sonarcloud.io) - Pushing covered Go apps to Dockerhub 

![](https://assets-eu-01.kc-usercontent.com/45f00125-6dea-0121-0efe-ce8937882537/7fd02d4d-94bc-4d88-a77e-db2d6e81a658/body-3cedacd4-0d81-4da8-986c-7978d6765e79_SC-SQ%2BBlog%2B-%2Bbranding%25402x.png?w=1432&h=448&auto=format&fit=crop)

## Overview 📌
This project is a pratical example of how to create a Continuous Integration (CI) to perform code coverage, analysis and quality checks automatically with [Sonarcloud](https://sonarcloud.io) & Github Actions. This integration enables you to automatically run code analysis using SonarCloud and to provide feedback on the quality of your code directly in your GitHub repository.

For this project we are deploying an illustrative example Go app through CI/CD workflows to Dockerhub.

## Continuous Integration (CI) Workflow (GitActions)

* :open_file_folder: Checkout the code from your repository
* :hammer_and_wrench: Set up your build environment (such as installing dependencies)
* :microscope: Run your build and test scripts
* :mag: Run the SonarCloud scanner to analyze your code
* :chart_with_upwards_trend: Publish the analysis results to SonarCloud

[Check the ci.yml file](https://github.com/kevencript/sonarcloud-ci-example/blob/main/.github/workflows/ci.yaml)

# Complete flow example

In this section we will exemplify the whole process (Creating a PR & checking the analysis result) with images

## 1 - Create a PR with code changes :octocat:

* For this example we will change the Hashes (we have go function and tests to verify these hashess into the main go example. Check the file for a better understanding) as a plan to trigger the code coverage analysis and exemplify how Sonarqube can be powerful code analysis tool [(Check the PR)](https://github.com/kevencript/sonarcloud-ci-example/pull/4/files)
> It's  important to comprehend that the Sonarcloud coverage analysis will be executed only into the changed lines of the code, in order to have the percentage of coverage code based on changes.

## 2 - Wait for the CI flow 🕣

* Now that the PR was made, the Continuous Integration (CI) process will run and make the test coverage. After it, Sonarqube will scan it and provide a feedback related to bugs, security issues and coverage percent.

<img src="https://imgur.com/Uj8jfUg.png" width="850px" heigth="100%">

## 3 - Check for bot Sonarqube feedback comment ✅

* Once the CI workflow is finished, it will automatically generate a comment with some stat analysis from Sonarqube & code coverage
<img src="https://imgur.com/yagzZIo.png" width="850px" heigth="100%">
> Here we can see the  Bugs, Vulnerabilities, Security Hotspot & Code Smells
> From here, we are allowed to merge the code

### 4 - Checking the Analysis on Sonarcloud UI 💻

* Now that everything is tested and covered, we can check the analysis into the Sonarcloud UI and have access to much more information related to our code

<img src="https://imgur.com/qcXygDN.png" width="850px" heigth="100%">

# Sonarcloud Config 🔧

## What did i create on SonarCloud 

#### 🚀 Create a project and point to my Github repo
I set up a SonarCloud project for my GitHub repository using the SonarCloud UI. By integrating SonarCloud with my GitHub repository, I can easily monitor the code quality and track the progress of code improvements over time.

#### 🔧 Define "GitActions" as main strategy
To automate my continuous integration pipeline, I used GitHub Actions as my main strategy. This allowed me to run SonarCloud code analysis on every push to my main branch. GitHub Actions is a powerful tool for automating repetitive tasks, and it integrated seamlessly with SonarCloud.

#### 🎛️ Configure Quality Profiles, Rules, and Quality Gates

In addition to setting up my SonarCloud project and integrating it with my GitHub repository using GitHub Actions, I also configured the Quality Profiles, Rules, and Quality Gates in SonarCloud. This allowed me to tailor the code analysis to my specific needs and ensure that I am maintaining the highest code quality standards. By setting up Quality Gates, I can easily track and ensure that my code meets the predefined criteria for quality.

#### 🔒 Create the SONAR_TOKEN secret into Actions Secrets
To ensure the security of my SonarCloud integration, I added the SONAR_TOKEN to my CI secrets. This token ensures that only authorized users can access and use the SonarCloud API, protecting my repository from potential security risks.

Overall, using GitHub Actions as my main strategy for automating my CI/CD pipeline and integrating with SonarCloud has been an effective way to improve my code quality and streamline my development process. By catching and resolving issues early on, I can ensure that my code is always of the highest quality, and I can focus on building high-quality applications that meet the needs of my users (example case).

---

# Sonarqube Overview 📝

## What is "Sonarqube"? 

[SonarQube](https://sonarcloud.io) is an open-source platform for continuous code quality inspection and analysis. It provides detailed reports on code issues and allows developers to track and improve the quality of their code over time. With SonarQube, developers can quickly identify and fix issues in their code, improving the maintainability and reliability of their applications. Additionally, it can be integrated with popular development tools for automatic code analysis and instant feedback to developers.

## Sonarqube usual platform therms
Here we have some concerns related to how we can create rules in which will define what is a bug or something critical, quality profiles and the quality gates, in which will define the quality standarts for our analysis.

### Sonarqube: Quality Profiles 👥
Quality Profiles are a key component of SonarCloud that allows you to define a set of rules to enforce code quality standards across your projects. You can choose from a variety of pre-defined profiles or create your own custom profile. The profile can be applied to individual projects or across your organization.

### Sonarqube: Rules 📜
SonarCloud provides a set of pre-defined rules for different programming languages. These rules check for various code quality issues, such as complexity, duplications, and security vulnerabilities. You can also create custom rules to meet specific requirements. The rules are applied to the code during analysis and reported in the SonarCloud dashboard.

### Quality Gates 🚦
Quality Gates are used to define the quality standards that must be met before code can be deployed to production. The Quality Gates in SonarCloud can be customized to include specific quality checks, such as code coverage, maintainability, and security. If the code does not meet the quality standards defined in the Quality Gate, the build will fail, and the issues will be reported in the SonarCloud dashboard.

Overall, these three components work together to ensure that your code meets the highest quality standards and helps you identify potential issues before they become a problem. With SonarCloud, you can create custom Quality Profiles, define rules to enforce standards, and set Quality Gates to ensure that your code is ready for deployment.






