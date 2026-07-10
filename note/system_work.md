**Multi-Tenant SaaS Mess Management System**। অর্থাৎ একটি অ্যাপ ব্যবহার করে অনেকগুলো মেস/হোস্টেল (Tenant) আলাদা আলাদাভাবে তাদের নিজস্ব ডাটা পরিচালনা করবে।

---

# পুরো সিস্টেম কীভাবে কাজ করবে

এখানে মূলত ৩ ধরনের Actor থাকবে।

```
Super Admin (System Owner)
        │
        ▼
Tenant (Mess Owner)
        │
        ▼
Users (Manager, Member, Cook, Accountant)
```

---

# ১. Super Admin

এটি অ্যাপের মালিক। একজন বা খুব কম সংখ্যক থাকবে।

Super Admin-এর কাজ:

* নতুন Tenant তৈরি করা
* Subscription Plan তৈরি করা
* Role ও Permission তৈরি করা
* Tenant Active/Inactive করা
* Payment দেখা
* Invoice দেখা
* Audit Log দেখা

Super Admin সকল Tenant-এর উপর নিয়ন্ত্রণ রাখে।

---

# ২. Tenant (Mess Owner)

ধরো,

```
Rahim Mess
```

এর Owner Registration করল।

তখন `TENANTS` টেবিলে একটি রেকর্ড হবে।

| Field         | Value                                     |
| ------------- | ----------------------------------------- |
| tenant_name   | Rahim Mess                                |
| tenant_code   | RM001                                     |
| sub_domain    | rahimmess                                 |
| database_name | rahim_db                                  |
| email         | [owner@gmail.com](mailto:owner@gmail.com) |

এখন Owner Login করবে

```
https://rahimmess.messmanager.com
```

---

# ৩. User

Tenant তৈরি হওয়ার পরে Owner বিভিন্ন User তৈরি করবে।

যেমন

```
Manager
Cook
Accountant
Member
```

সব User `USERS` টেবিলে থাকবে।

---

# Role কে তৈরি করবে?

তোমার ERD অনুযায়ী

```
ROLES
```

সবার জন্য Common।

এগুলো Super Admin একবার তৈরি করবে।

যেমন

```
Super Admin
Tenant Admin
Manager
Accountant
Cook
Member
```

Tenant এগুলো ব্যবহার করবে।

---

# Permission কে তৈরি করবে?

Super Admin

```
Create Meal
Edit Meal
Delete Meal

Create Deposit
Edit Deposit

Create Bazar
Delete Bazar

View Report
```

এগুলো

```
PERMISSIONS
```

টেবিলে থাকবে।

---

# Role Permission

```
Manager

↓

Create Meal ✔

↓

Create Deposit ✔

↓

Delete User ✘
```

এই Mapping থাকবে

```
ROLE_PERMISSIONS
```

টেবিলে।

---

# Tenant

```
TENANTS
```

এটি পুরো একটি Organization।

যেমন

```
City Mess

Green Villa

Hostel One
```

প্রতিটি Tenant-এর

* আলাদা User
* আলাদা Meal
* আলাদা Deposit
* আলাদা Bazar

থাকবে।

---

# Users

প্রতিটি User একটি Tenant-এর সদস্য।

```
Rahim Mess

↓

Manager

↓

Accountant

↓

Cook

↓

Members
```

সবাই Users Table-এ থাকবে।

---

# Plans

ধরো

```
Free

Basic

Premium

Enterprise
```

এগুলো

```
PLANS
```

টেবিলে থাকবে।

Super Admin এগুলো তৈরি করবে।

---

# Subscription

যখন একটি Tenant

```
Premium Plan
```

কিনবে।

তখন

```
SUBSCRIPTIONS
```

এ রেকর্ড হবে।

যেমন

```
Rahim Mess

↓

Premium

↓

1 Year
```

---

# Payment

Tenant যখন Payment করবে

```
5000 Tk
```

তখন

```
PAYMENTS
```

এ রেকর্ড হবে।

---

# Invoice

Payment Success হলে

```
Invoice Generate
```

হবে।

```
INV-2026001
```

---

# Meals

প্রতিদিন Member-এর Meal Entry।

যেমন

| User  | Breakfast | Lunch | Dinner |
| ----- | --------- | ----- | ------ |
| Rahim | 1         | 1     | 1      |
| Karim | 0         | 1     | 1      |

সব

```
MEALS
```

টেবিলে যাবে।

---

# Deposits

Member টাকা জমা দিলে

```
1000 Tk
```

```
DEPOSITS
```

টেবিলে যাবে।

---

# Bazars

আজ বাজার হয়েছে

```
Chicken

Rice

Oil
```

প্রথমে

```
BAZARS
```

এ Header থাকবে।

যেমন

```
Friday Market
```

---

# Bazar Items

তারপর Item গুলো

```
BAZAR_ITEMS
```

এ যাবে।

| Item    | Qty  | Price |
| ------- | ---- | ----- |
| Rice    | 20kg | 1200  |
| Chicken | 10kg | 2500  |
| Oil     | 5L   | 800   |

---

# Audit Logs

যখন কেউ

```
Delete Meal

Update Deposit

Create User
```

করবে।

সব Log

```
AUDIT_LOGS
```

এ Save হবে।

Example

```
User:

Manager

Action:

Delete Meal

Time:

10:30 AM
```

---

# Notifications

যখন

```
Subscription Expiring

Meal Added

Deposit Received

New User Added
```

হবে।

তখন

```
NOTIFICATIONS
```

এ Save হবে।

---

# Login Flow

```
User

↓

citymess.messmanager.com

↓

Echo Middleware

↓

Extract

citymess

↓

TENANTS Table

↓

tenant_id = 5

↓

User Login

↓

USERS

WHERE tenant_id=5

↓

Dashboard
```

---

# Database Flow

```
Super Admin
        │
        ▼
Create Tenant
        │
        ▼
TENANTS
        │
        ▼
Assign Plan
        │
        ▼
SUBSCRIPTIONS
        │
        ▼
Tenant Login
        │
        ▼
Create Users
        │
        ▼
Assign Roles
        │
        ▼
Daily Meal Entry
        │
        ▼
Deposit Entry
        │
        ▼
Bazar Entry
        │
        ▼
Reports
```

---

# এই ERD-তে আমি যে Workflow সাজেস্ট করি

1. **Super Admin**

   * Plan তৈরি করবে।
   * Role তৈরি করবে।
   * Permission তৈরি করবে।
   * নতুন Tenant তৈরি বা অনুমোদন করবে।
   * Subscription ও Payment পর্যবেক্ষণ করবে।

2. **Tenant Admin (Mess Owner)**

   * নিজের Tenant-এর User তৈরি করবে।
   * User-দের Role Assign করবে।
   * Meal, Deposit, Bazar, Notification পরিচালনা করবে।
   * শুধুমাত্র নিজের Tenant-এর ডাটা দেখতে পারবে।

3. **Manager/Accountant/Cook/Member**

   * তাদের Role অনুযায়ী Permission পাবে।
   * Manager Meal ও Bazar পরিচালনা করতে পারে।
   * Accountant Deposit ও হিসাব পরিচালনা করতে পারে।
   * Member নিজের Meal, Deposit ও Notification দেখতে পারে।

এই ডিজাইনটি একটি **shared-database, multi-tenant SaaS architecture**-এর জন্য ভালো ভিত্তি। ভবিষ্যতে এতে **Reports**, **Expenses**, **Monthly Bill Generation**, **Meal Rate Calculation**, **Balance Sheet**, **OTP Login**, এবং **JWT Authentication** যোগ করা সহজ হবে।
