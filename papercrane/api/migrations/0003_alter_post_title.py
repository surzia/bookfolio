# Generated by Django 4.0.1 on 2022-08-21 02:32

import datetime
from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('api', '0002_alter_post_created_alter_post_updated'),
    ]

    operations = [
        migrations.AlterField(
            model_name='post',
            name='title',
            field=models.CharField(default=datetime.date(2022, 8, 21), max_length=50, unique=True, verbose_name='标题'),
        ),
    ]
